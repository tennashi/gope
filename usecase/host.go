package usecase

import (
	"log"

	"github.com/tennashi/gope/domain/repository"

	"github.com/tennashi/gope/domain"
)

// Host is usecase
type Host struct {
	hr repository.Hosts
}

func NewHost(hr repository.Hosts) *Host {
	return &Host{hr}
}

// List show hosts file
func (h *Host) List(hostsFile string) (hosts *domain.Hosts, err error) {
	hosts, err = h.hr.Get(hostsFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

// HostsList show hosts file list
func (h *Host) HostsList(hostsDir string) (hostsNames []string, err error) {
	return
}
