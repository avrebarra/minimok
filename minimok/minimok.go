package minimok

import (
	"net/http"

	"github.com/avrebarra/minimok/mokserver"
	"gopkg.in/yaml.v2"
)

type Handler struct {
	MokSpec mokserver.Spec
	Handler http.Handler
}

type ConfigFile struct {
	MokSpecs []mokserver.Spec `yaml:"minimok"`
}

func ParseConfigFile(bits []byte) (cfg ConfigFile, err error) {
	err = yaml.Unmarshal(bits, &cfg)
	if err != nil {
		return
	}

	return
}
