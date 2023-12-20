package examine_repository

import "github.com/kaiiorg/go-bif-examine/pkg/models"

type ExamineRepository interface {
	GetAllProjects() ([]*models.Project, error)
	CreateProject(project *models.Project) (uint, error)
	DeleteProject(projectId uint) error

	CreateManyBifs(bifs []*models.Bif) error

	CreateManyResources(resources []*models.Resource) error
}
