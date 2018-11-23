package repository

import (
	"github.com/tennashi/gope/domain"
)

// Hosts is repository
type Hosts interface {
	GetHostFiles(dirPath string) ([]*domain.HostFile, error)
	GetHosts(fileName string) (*domain.Hosts, error)
	GetHost(fileName, hostName string) (*domain.Host, error)
}
