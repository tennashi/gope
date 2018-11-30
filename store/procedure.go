package store

import (
	"io/ioutil"
	"path/filepath"

	toml "github.com/pelletier/go-toml"
	"github.com/tennashi/gope/domain"
)

// Procedure implements Procedure repository.
type Procedure struct{}

// GetProcedureFiles returns procedure files.
func (p *Procedure) GetProcedureFiles(basePath string) ([]*domain.ProcedureFile, error) {
	glob := filepath.Join(basePath, "procedures", "*.gope")
	fileNames, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	proFiles := make([]*domain.ProcedureFile, len(fileNames))
	for i, v := range fileNames {
		proFiles[i] = &domain.ProcedureFile{
			Name: getName(v),
			Path: v,
		}
	}
	return proFiles, nil
}

// GetProcedureFile returns a procedure file.
func (p *Procedure) GetProcedureFile(basePath, name string) (*domain.ProcedureFile, error) {
	path := filepath.Join(basePath, "procedures", name+".gope")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &domain.ProcedureFile{
		Name:  name,
		Path:  path,
		Bytes: bytes,
	}, nil
}

// GetProcedure returns a procedure.
func (p *Procedure) GetProcedure(proFile *domain.ProcedureFile) (procedure *domain.Procedure, err error) {
	procedure = &domain.Procedure{}
	if err = toml.Unmarshal(proFile.Bytes, procedure); err != nil {
		return nil, err
	}
	return
}
