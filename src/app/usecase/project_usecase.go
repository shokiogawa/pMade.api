package usecase

import (
	"fmt"
	"pMade.api/src/app/domain/model"
	"pMade.api/src/app/domain/repository"
)

type ProjectUseCase interface {
	GetProject() ([]*model.Project, error)
}

type projectUseCase struct {
	projectRepository repository.ProjectRepository
}

func (pu *projectUseCase) GetProject() (projects []*model.Project, err error) {
	projects, err = pu.projectRepository.GetProject()
	if err != nil {
		fmt.Println(err)
	}
	return projects, err
}

func NewProjectUseCase(pr repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{
		projectRepository: pr,
	}
}
