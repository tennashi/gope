package handler

import (
	"log"
	"net/http"
	"path/filepath"

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

// GetHosts show Hosts list
func (h *Host) GetHosts(c *context.Context) error {
	fileName := c.Param("name") + ".toml"
	filePath := filepath.Join(c.Config.GetString("base_path"), "hosts", fileName)
	ret, err := h.uc.List(filePath)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, ret)
}
