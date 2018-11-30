package context

import (
	"github.com/labstack/echo"
	"github.com/tennashi/gope/domain"
)

// Project represents the project context.
type Project struct {
	echo.Context
	Project *domain.Project
}

// NewProjectContext constructs a ner project context.
func NewProjectContext(ec echo.Context, project *domain.Project) *Project {
	return &Project{ec, project}
}
