package examine_repository

import (
	"errors"

	"github.com/kaiiorg/go-bif-examine/pkg/models"

	"gorm.io/gorm"
)

func (r *GormExamineRepository) GetAllProjects() ([]*models.Project, error) {
	results := []*models.Project{}
	err := r.db.Where("deleted_at IS NULL").Find(&results).Error
	if err != nil {
		r.log.Error().Err(err).Msg("Failed to get all projects")
		return nil, err
	}
	return results, nil
}

func (r *GormExamineRepository) CreateProject(project *models.Project) (uint, error) {
	if err := r.db.Create(project).Error; err != nil {
		return 0, err
	}
	return project.ID, nil
}

func (r *GormExamineRepository) DeleteProject(projectId uint) error {
	project := &models.Project{
		Model: gorm.Model{ID: projectId},
	}
	err := r.db.Delete(project).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
