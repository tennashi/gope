package handler

import (
	"net/http"

	"github.com/tennashi/gope/store"

	"github.com/labstack/echo"
	"github.com/tennashi/gope/usecase"

	"github.com/tennashi/gope/context"
)

// Procedure represents a procedure handler.
type Procedure struct {
	uc *usecase.Procedure
}

// NewProcedure construct a new procedure handler.
func NewProcedure() *Procedure {
	uc := usecase.NewProcedure(&store.Procedure{})
	return &Procedure{uc}
}

// GetProcedureFiles show procedure files.
func (p *Procedure) GetProcedureFiles(c *context.Project) error {
	basePath := c.Project.BasePath
	ret, err := p.uc.GetProcedureFiles(basePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}

// GetProcedure show the procedure.
func (p *Procedure) GetProcedure(c *context.Project) error {
	name := c.Param("name")
	basePath := c.Project.BasePath
	ret, err := p.uc.GetProcedure(basePath, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}
