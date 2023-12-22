package examine_repository

import (
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/models"
)

func (r *GormExamineRepository) CreateManyResources(resources []*models.Resource) error {
	batchSize := 100
	resourcesChan := make(chan *models.Resource, 5)
	batchChan := make(chan []*models.Resource, 5)
	errChan := make(chan error, 5)
	wg := &sync.WaitGroup{}

	go r.batchResources(batchSize, resourcesChan, batchChan, wg)
	go r.insertBatches(batchChan, errChan, wg)

	// Loop through the resources, make sure their BifID and ProjectID are set, then pass them onto the batcher go routine
	for _, resource := range resources {
		resource.BifID = resource.Bif.ID
		resource.ProjectID = resource.Project.ID
		resourcesChan <- resource
	}
	close(resourcesChan)

	errCount := 0
	var finalErr error

	for err := range errChan {
		// Don't add more than 5 errors to the finalErr
		if errCount > 5 {
			continue
		}
		finalErr = errors.Join(finalErr, err)
		errCount++
	}

	wg.Wait()
	return finalErr
}

func (r *GormExamineRepository) FindProjectResourcesForBif(projectId uint, bifId uint) ([]*models.Resource, error) {
	resources := []*models.Resource{}
	err := r.db.Where("project_id = ? AND bif_id = ? AND deleted_at IS NULL", projectId, bifId).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (r *GormExamineRepository) GetResourceById(resourceId uint) (*models.Resource, error) {
	resource := &models.Resource{
		Model: gorm.Model{
			ID: resourceId,
		},
	}
	err := r.db.First(resource).Error
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (r *GormExamineRepository) GetResourceForWhisper() (*models.Resource, error) {
	resource := &models.Resource{}

	// TODO find a better way of doing this; there's a small chance that we could query for a record,
	// then while we're about to update the job started column, another request grabs the same record
	err := r.db.
		Where("deleted_at IS NULL AND offset_to_data != 0 AND size != 0 AND job_duration IS NULL").
		Order("job_started").
		Limit(1).
		First(resource).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Model(resource).
		Updates(map[string]interface{}{"job_started": time.Now()}).
		Error
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (r *GormExamineRepository) UpdateResource(resource *models.Resource) error {
	return r.db.Save(resource).Error
}
