package repository

import (
	"github.com/tennashi/gope/domain"
)

// Procedure defines the CRUD interface of the procedure model.
type Procedure interface {
	GetProcedureFiles(basePath string) ([]*domain.ProcedureFile, error)
	GetProcedureFile(basePath, name string) (*domain.ProcedureFile, error)
	GetProcedure(proFile *domain.ProcedureFile) (*domain.Procedure, error)
}
