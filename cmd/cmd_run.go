package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/avrebarra/minimok/mux"
	"github.com/facebookgo/grace/gracehttp"
	"gopkg.in/go-playground/validator.v9"
)

type ConfigCommandStart struct {
	Port       int    `validate:"required"`
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
	// load config file
	fmt.Printf("using configfile %s\n", c.config.ConfigPath)
	bits, err := ioutil.ReadFile(c.config.ConfigPath)
	if err != nil {
		return
	}
	spec, err := mux.MuxSpecFromYAML(bits)
	if err != nil {
		return
	}

	// setup mux
	m := mux.New()
	h, err := m.GetHandler(context.Background())
	if err != nil {
		err = fmt.Errorf("failed getting initial handler")
		return
	}

	// apply spec
	err = m.ApplySpec(context.Background(), spec)
	if err != nil {
		return
	}

	// start server
	fmt.Printf("starting up server on http://localhost:%d\n", c.config.Port)
	err = gracehttp.Serve(&http.Server{Addr: fmt.Sprint(":", c.config.Port), Handler: h})
	if err != nil {
		err = fmt.Errorf("failed starting server")
		os.Exit(1)
	}

	return
}
