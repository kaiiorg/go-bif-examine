package examine_repository

import (
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/models"
)

func (r *GormExamineRepository) CreateManyResources(resources []*models.Resource) error {
	batchSize := 100
	errGroup := errgroup.Group{}
	resourcesChan := make(chan *models.Resource, 5)
	batchChan := make(chan []*models.Resource, 5)

	errGroup.Go(func() error {
		return r.batchResources(batchSize, resourcesChan, batchChan)
	})
	errGroup.Go(func() error {
		return r.insertBatches(batchChan)
	})
	errGroup.Go(func() error {
		// Loop through the resources, make sure their BifID and ProjectID are set, then pass them onto the batcher go routine
		for _, resource := range resources {
			resource.BifID = resource.Bif.ID
			resource.ProjectID = resource.Project.ID
			resourcesChan <- resource
		}
		close(resourcesChan)
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return err
	}

	return nil
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
		Where("deleted_at IS NULL AND offset_to_data != 0 AND size != 0 AND job_duration = ''").
		Order("job_started DESC").
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
