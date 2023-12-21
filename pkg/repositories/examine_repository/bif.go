package examine_repository

import "github.com/kaiiorg/go-bif-examine/pkg/models"

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

func (r *GormExamineRepository) UpdateBif(bif *models.Bif) error {
	return r.db.Save(bif).Error
}
