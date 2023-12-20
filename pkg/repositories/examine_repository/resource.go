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

// batchResources takes resources from resourcesChan and batches them into batches of batchSize and then sends them to
// batchChan. When the resourcesChan closes, the remaining batch is sent if it has at least one resource in it.
// TODO generalize this with generics?
func (r *GormExamineRepository) batchResources(batchSize int, resourcesChan chan *models.Resource, batchChan chan []*models.Resource, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	batch := []*models.Resource{}
	for resource := range resourcesChan {
		batch = append(batch, resource)
		if len(batch) >= batchSize {
			batchChan <- batch
			batch = []*models.Resource{}
		}
	}
	if len(batch) > 0 {
		batchChan <- batch
	}
	close(batchChan)
}

// insertBatches will take the batches from batchChan and insert them into the database, sending any error via the errorsChan
func (r *GormExamineRepository) insertBatches(batchChan chan []*models.Resource, errorsChan chan error, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for batch := range batchChan {
		err := r.db.Create(batch).Error
		if err != nil {
			errorsChan <- err
		}
	}
	close(errorsChan)
}
