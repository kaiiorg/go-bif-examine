package examine_repository

import (
	"github.com/kaiiorg/go-bif-examine/pkg/models"
)

// batchResources takes resources from resourcesChan and batches them into batches of batchSize and then sends them to
// batchChan. When the resourcesChan closes, the remaining batch is sent if it has at least one resource in it.
// TODO generalize this with generics?
func (r *GormExamineRepository) batchResources(batchSize int, resourcesChan chan *models.Resource, batchChan chan []*models.Resource) error {
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
	return nil
}

// insertBatches will take the batches from batchChan and insert them into the database, sending any error via the errorsChan
func (r *GormExamineRepository) insertBatches(batchChan chan []*models.Resource) error {
	for batch := range batchChan {
		err := r.db.Create(batch).Error
		if err != nil {
			return err
		}
	}
	return nil
}
