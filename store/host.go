package store

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
	"github.com/tennashi/gope/domain"
)

type Hosts struct{}

func (t *Hosts) Get(fileName string) (hosts *domain.Hosts, err error) {
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
