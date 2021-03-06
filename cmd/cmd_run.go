package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/avrebarra/minimok/minimok"
	"github.com/avrebarra/minimok/mokserver"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

type configStart struct {
	ConfigPath string `validate:"required"`
}

func runStart(cfg configStart) (err error) {
	type ConfigFile struct {
		MokSpecs []mokserver.Spec `yaml:"minimok"`
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// validate config
	if err := validator.New().Struct(cfg); err != nil {
		return err
	}

	// cancel context on interrupt
	exitc := make(chan os.Signal, 1)
	signal.Notify(exitc, os.Interrupt)
	go func() {
		select {
		case <-exitc:
			cancel()
		}
	}()

	// load config file
	spec := ConfigFile{}
	fmt.Printf("using configfile %s\n", cfg.ConfigPath)
	bits, err := ioutil.ReadFile(cfg.ConfigPath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(bits, &spec)
	if err != nil {
		return
	}

	// setup minimok
	mok, err := minimok.New(minimok.Config{
		MokSpecs: spec.MokSpecs,
	})
	if err != nil {
		return
	}

	// start server(s)
	err = mok.Start(ctx)
	if err != nil {
		return
	}

	return
}
