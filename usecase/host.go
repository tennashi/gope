package usecase

import (
	"log"

	"github.com/tennashi/gope/domain/repository"

	"github.com/tennashi/gope/domain"
)

// Host represents a host usecase.
type Host struct {
	hr repository.Hosts
}

// NewHost constructs a host usecase.
func NewHost(hr repository.Hosts) *Host {
	return &Host{hr}
}

// GetHostFiles uses host repository to return host files.
func (h *Host) GetHostFiles(basePath string) (hostFiles []*domain.HostFile, err error) {
	hostFiles, err = h.hr.GetHostFiles(basePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

// GetHosts uses host repository to return hosts.
func (h *Host) GetHosts(basePath, name string) (hosts *domain.Hosts, err error) {
	hostFile, err := h.hr.GetHostFile(basePath, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	hosts, err = h.hr.GetHosts(hostFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

// GetHost uses host repository to return a host.
func (h *Host) GetHost(basePath, fileName, hostName string) (host *domain.Host, err error) {
	hostFile, err := h.hr.GetHostFile(basePath, fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	host, err = h.hr.GetHost(hostFile, hostName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
