package cmd

import (
	"gopkg.in/go-playground/validator.v9"
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
	return
}
