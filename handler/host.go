package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tennashi/gope/context"
	"github.com/tennashi/gope/store"
	"github.com/tennashi/gope/usecase"
)

// Host represent a host handler.
type Host struct {
	uc *usecase.Host
}

// NewHost construct a new host handler.
func NewHost() *Host {
	repo := &store.Hosts{}
	uc := usecase.NewHost(repo)
	return &Host{uc}
}

// GetHostFiles show host files.
func (h *Host) GetHostFiles(c *context.Project) error {
	basePath := c.Project.BasePath
	ret, err := h.uc.GetHostFiles(basePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}

// GetHosts show the host list.
func (h *Host) GetHosts(c *context.Project) error {
	name := c.Param("name")
	basePath := c.Project.BasePath
	ret, err := h.uc.GetHosts(basePath, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}

// GetHost show the host.
func (h *Host) GetHost(c *context.Project) error {
	fileName := c.Param("name")
	hostName := c.Param("host")
	basePath := c.Project.BasePath
	ret, err := h.uc.GetHost(basePath, fileName, hostName)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}
