package store

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	toml "github.com/pelletier/go-toml"
	"github.com/tennashi/gope/domain"
)

// Hosts implement Hosts repository
type Hosts struct{}

func (h *Hosts) GetHostFiles(dirPath string) ([]*domain.HostFile, error) {
	glob := filepath.Join(dirPath, "*.toml")
	fileNames, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	hostFiles := make([]*domain.HostFile, len(fileNames))
	for i, v := range fileNames {
		hostFiles[i] = &domain.HostFile{
			Name: getName(v),
			Path: v,
		}
	}
	return hostFiles, nil
}

func (h *Hosts) GetHosts(fileName string) (hosts *domain.Hosts, err error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	hosts = &domain.Hosts{}
	if err = toml.Unmarshal(bytes, hosts); err != nil {
		return nil, err
	}
	return
}

func (h *Hosts) GetHost(fileName, hostName string) (host *domain.Host, err error) {
	hosts, err := h.GetHosts(fileName)
	if err != nil {
		return nil, err
	}
	for _, v := range hosts.Hosts {
		if v.Name == hostName {
			host = &v
		}
	}
	if host == nil {
		return nil, fmt.Errorf("%s is not found in %s", hostName, fileName)
	}
	return
}

func getName(filePath string) string {
	_, fileName := filepath.Split(filePath)
	name := strings.TrimSuffix(fileName, ".toml")
	return name
}
