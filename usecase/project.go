package usecase

import (
	"log"

	"github.com/tennashi/gope/domain"
	"github.com/tennashi/gope/domain/repository"
)

// Project represents a project usecase.
type Project struct {
	pr repository.Project
}

// NewProject constructs a project usecase.
func NewProject(pr repository.Project) *Project {
	return &Project{pr}
}

// GetProject uses project repository to return a project.
func (p *Project) GetProject(name string) (project *domain.Project, err error) {
	project, err = p.pr.GetProject(name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
