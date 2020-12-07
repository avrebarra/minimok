package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/avrebarra/minimok/minimok"
	"github.com/avrebarra/minimok/mux"
	"github.com/facebookgo/grace/gracehttp"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v3"
)

type ConfigCommandStart struct {
	ConfigPath string `validate:"required"`
}

type CommandStart struct {
	config ConfigCommandStart
}

func NewCommandStart(cfg ConfigCommandStart) CommandStart {
	if err := validator.New().Struct(cfg); err != nil {
		panic(err)
	}

	cmd := CommandStart{config: cfg}

	return cmd
}

func (c *CommandStart) Run() (err error) {
	ctx := context.Background()

	// load config file
	fmt.Printf("using configfile %s\n", c.config.ConfigPath)
	bits, err := ioutil.ReadFile(c.config.ConfigPath)
	if err != nil {
		return
	}
	spec, err := c.parseConfigFile(bits)
	if err != nil {
		return
	}

	// setup minimok
	mok := minimok.New(minimok.Config{
		MuxSpecs: spec.MuxSpecs,
	})

	// get minimok handlers
	handlers, err := mok.GetHandlers(ctx)
	if err != nil {
		return
	}

	// start server(s)
	for _, ms := range handlers {
		portaddr := ms.MuxSpec.Port
		fmt.Printf("starting up mokserver %s on http://localhost:%d\n", ms.MuxSpec.Name, portaddr)
		err = gracehttp.Serve(&http.Server{Addr: fmt.Sprint(":", portaddr), Handler: ms.Handler})
		if err != nil {
			err = fmt.Errorf("failed starting server: %s", ms.MuxSpec.Name)
			os.Exit(1)
		}
	}

	return
}

// ***

type CommandStartConfigFile struct {
	MuxSpecs []mux.MuxSpec `yaml:"minimok"`
}

func (c *CommandStart) parseConfigFile(bits []byte) (cfg CommandStartConfigFile, err error) {
	err = yaml.Unmarshal(bits, &cfg)
	if err != nil {
		return
	}

	return
}
