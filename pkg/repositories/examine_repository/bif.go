package examine_repository

import (
	"github.com/kaiiorg/go-bif-examine/pkg/models"

	"gorm.io/gorm"
)

func (r *GormExamineRepository) CreateManyBifs(bifs []*models.Bif) error {
	return r.db.Create(&bifs).Error
}

func (r *GormExamineRepository) GetBifByNormalizedNameOrNameInKey(normalizedBifName, bifNameInKey string) (*models.Bif, error) {
	bifRecord := &models.Bif{}
	var err error

	if normalizedBifName == "" {
		err = r.db.Where("name_in_key = ? AND deleted_at IS NULL", bifNameInKey).First(&bifRecord).Error
	} else {
		err = r.db.Where("name = ? AND deleted_at IS NULL", normalizedBifName).First(&bifRecord).Error
	}

	if err != nil {
		return nil, err
	}

	return bifRecord, nil
}

func (r *GormExamineRepository) GetBifById(bifId uint) (*models.Bif, error) {
	bif := &models.Bif{
		Model: gorm.Model{
			ID: bifId,
		},
	}
	err := r.db.First(bif).Error
	if err != nil {
		return nil, err
	}
	return bif, nil
}

func (r *GormExamineRepository) GetBifsMissingContent(projectId uint) ([]*models.Bif, error) {
	bifs := []*models.Bif{}
	r.db.
		Where("object_key IS NULL").
		Where(
			"id IN (?)",
			r.db.Model(models.Resource{}).
				Select("DISTINCT bif_id").
				Where("project_id = ?", projectId),
		).
		Find(&bifs)
	return bifs, nil
}

func (r *GormExamineRepository) UpdateBif(bif *models.Bif) error {
	return r.db.Save(bif).Error
}
