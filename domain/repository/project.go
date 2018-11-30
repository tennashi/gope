package repository

import "github.com/tennashi/gope/domain"

// Project defines the CRUD interface of the project model.
type Project interface {
	GetProject(name string) (*domain.Project, error)
}
