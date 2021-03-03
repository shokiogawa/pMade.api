package repository

import (
	"pMade.api/src/app/domain/model"
)

type ProjectRepository interface {
	GetProject() ([]*model.Project, error)
}
