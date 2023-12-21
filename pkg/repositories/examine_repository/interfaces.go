package examine_repository

import "github.com/kaiiorg/go-bif-examine/pkg/models"

type ExamineRepository interface {
	GetAllProjects() ([]*models.Project, error)
	GetProjectById(projectId uint) (*models.Project, error)
	CreateProject(project *models.Project) (uint, error)
	DeleteProject(projectId uint) error

	CreateManyBifs(bifs []*models.Bif) error
	GetBifByNormalizedNameOrNameInKey(normalizedBifName, bifNameInKey string) (*models.Bif, error)
	GetBifById(bifId uint) (*models.Bif, error)
	UpdateBif(bif *models.Bif) error

	CreateManyResources(resources []*models.Resource) error
	FindProjectResourcesForBif(projectId uint, bifId uint) ([]*models.Resource, error)
	GetResourceById(resourceId uint) (*models.Resource, error)
	UpdateResource(resource *models.Resource) error
}
