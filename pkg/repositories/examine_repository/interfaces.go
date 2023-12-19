package examine_repository

import "github.com/kaiiorg/go-bif-examine/pkg/models"

type ExamineRepository interface {
	GetAllProjects() ([]*models.Project, error)
	DeleteProject(projectId uint) error
}
