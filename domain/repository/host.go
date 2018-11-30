package repository

import (
	"github.com/tennashi/gope/domain"
)

// Hosts defines the CRUD interface of the host model.
type Hosts interface {
	GetHostFiles(basePath string) ([]*domain.HostFile, error)
	GetHostFile(basePath, name string) (*domain.HostFile, error)
	GetHosts(hostFile *domain.HostFile) (*domain.Hosts, error)
	GetHost(hostFile *domain.HostFile, name string) (*domain.Host, error)
}
