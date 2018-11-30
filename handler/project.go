package handler

import (
	"net/http"

	"github.com/tennashi/gope/context"
	"github.com/tennashi/gope/domain/repository"

	"github.com/labstack/echo"
	"github.com/tennashi/gope/usecase"
)

// Project represent a project handler.
type Project struct {
	uc *usecase.Project
}

// NewProject construct a new project handler.
func NewProject(repo repository.Project) *Project {
	uc := usecase.NewProject(repo)
	return &Project{uc}
}

// GetProject show the project.
func (p *Project) GetProject(c *context.Project) error {
	ret := c.Project
	return c.JSON(http.StatusOK, ret)
}

// UseProjectContext create a context.Project and returns echo.HandlerFunc.
// This implements echo.MiddlewareFunc.
func (p *Project) UseProjectContext(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		projectName := ec.Param("project_name")
		project, err := p.uc.GetProject(projectName)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		c := context.NewProjectContext(ec, project)
		return f(c)
	}
}
