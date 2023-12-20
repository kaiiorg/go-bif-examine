package examine_repository

import "github.com/kaiiorg/go-bif-examine/pkg/models"

func (r *GormExamineRepository) CreateManyBifs(bifs []*models.Bif) error {
	return r.db.Create(&bifs).Error
}
