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

func (h *Host) GetHostFiles(dirPath string) (hostFiles []*domain.HostFile, err error) {
	hostFiles, err = h.hr.GetHostFiles(dirPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (h *Host) GetHosts(hostsFile string) (hosts *domain.Hosts, err error) {
	hosts, err = h.hr.GetHosts(hostsFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (h *Host) GetHost(hostsFile, hostsName string) (host *domain.Host, err error) {
	host, err = h.hr.GetHost(hostsFile, hostsName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
