package store

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	toml "github.com/pelletier/go-toml"

	"github.com/tennashi/gope/domain"
)

// Hosts implements Hosts repository.
type Hosts struct{}

// GetHostFiles return host files.
func (h *Hosts) GetHostFiles(dirPath string) ([]*domain.HostFile, error) {
	glob := filepath.Join(dirPath, "hosts", "*.toml")
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

// GetHostFile return a host file.
func (h *Hosts) GetHostFile(basePath, name string) (*domain.HostFile, error) {
	path := filepath.Join(basePath, "hosts", name+".toml")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &domain.HostFile{
		Name:  name,
		Path:  path,
		Bytes: bytes,
	}, nil
}

// GetHosts return hosts.
func (h *Hosts) GetHosts(hostFile *domain.HostFile) (hosts *domain.Hosts, err error) {
	hosts = &domain.Hosts{}
	if err = toml.Unmarshal(hostFile.Bytes, hosts); err != nil {
		return nil, err
	}
	return
}

// GetHost return a host.
func (h *Hosts) GetHost(hostFile *domain.HostFile, name string) (*domain.Host, error) {
	hosts, err := h.GetHosts(hostFile)
	if err != nil {
		return nil, err
	}
	return findHost(hosts, name)
}

func findHost(hosts *domain.Hosts, name string) (host *domain.Host, err error) {
	for _, v := range hosts.Hosts {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("%s is not found in %s", name, hosts.Name)
}

func getName(filePath string) string {
	_, fileName := filepath.Split(filePath)
	name := strings.TrimSuffix(fileName, ".toml")
	return name
}
