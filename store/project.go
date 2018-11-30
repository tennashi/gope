package store

import (
	"fmt"
	"io/ioutil"

	toml "github.com/pelletier/go-toml"

	"github.com/tennashi/gope/domain"
)

// Project implements Project repository.
type Project struct {
	bytes []byte
}

// NewProject constructs a project.
func NewProject(filePath string) (*Project, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	project := &Project{bytes: bytes}
	return project, nil
}

// GetProject returns a project.
func (p *Project) GetProject(name string) (project *domain.Project, err error) {
	projects := &domain.Projects{}
	if err := toml.Unmarshal(p.bytes, projects); err != nil {
		return nil, err
	}
	for _, v := range projects.Projects {
		if v.Name == name {
			project = v
		}
	}
	if project == nil {
		return nil, fmt.Errorf("%s is not found", name)
	}

	return
}
