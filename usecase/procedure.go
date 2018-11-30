package usecase

import (
	"log"

	"github.com/tennashi/gope/domain/repository"

	"github.com/tennashi/gope/domain"
)

// Procedure represents a procedure usecase.
type Procedure struct {
	pr repository.Procedure
}

// NewProcedure constructs a procedure usecase.
func NewProcedure(pr repository.Procedure) *Procedure {
	return &Procedure{pr}
}

// GetProcedureFiles uses procedure repository to return procedure files.
func (p *Procedure) GetProcedureFiles(basePath string) (proFiles []*domain.ProcedureFile, err error) {
	proFiles, err = p.pr.GetProcedureFiles(basePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

// GetProcedure uses procedure repository to return a procedure.
func (p *Procedure) GetProcedure(basePath, name string) (procedure *domain.Procedure, err error) {
	proFile, err := p.pr.GetProcedureFile(basePath, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	procedure, err = p.pr.GetProcedure(proFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
