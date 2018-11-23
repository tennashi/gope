package handler

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"

	"github.com/tennashi/gope/context"
	"github.com/tennashi/gope/store"
	"github.com/tennashi/gope/usecase"
)

// Host represent host usecase
type Host struct {
	uc *usecase.Host
}

// NewHost is constructor
func NewHost() *Host {
	uc := usecase.NewHost(&store.Hosts{})
	return &Host{uc}
}

// GetHostFiles show host files
func (h *Host) GetHostFiles(c *context.Context) error {
	dirPath := filepath.Join(c.Config.GetString("base_path"), "hosts")
	ret, err := h.uc.GetHostFiles(dirPath)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}

// GetHosts show Hosts list
func (h *Host) GetHosts(c *context.Context) error {
	fileName := c.Param("name") + ".toml"
	filePath := filepath.Join(c.Config.GetString("base_path"), "hosts", fileName)
	ret, err := h.uc.GetHosts(filePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}

// GetHost show host
func (h *Host) GetHost(c *context.Context) error {
	fileName := c.Param("name") + ".toml"
	filePath := filepath.Join(c.Config.GetString("base_path"), "hosts", fileName)
	hostName := c.Param("host")
	ret, err := h.uc.GetHost(filePath, hostName)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, ret)
}
