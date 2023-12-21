package examine_repository

import (
	"errors"
	"sync"

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

func (r *GormExamineRepository) UpdateResource(resource *models.Resource) error {
	return r.db.Save(resource).Error
}
