package repository

import (
	"github.com/tennashi/gope/domain"
)

// Hosts is repository
type Hosts interface {
	Get(fileName string) (*domain.Hosts, error)
}
