package mux

import (
	"context"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
}

type Default struct {
	config Config
	spec   MuxSpec
	mux    *http.ServeMux
}

func New(cfg Config) Mux {
	if err := validator.New().Struct(cfg); err != nil {
		panic(err)
	}
	return &Default{
		config: cfg,
		mux:    http.NewServeMux(),
		spec:   MuxSpec{},
	}
}

func (e *Default) ApplySpec(ctx context.Context, spec MuxSpec) (err error) {
	e.spec = spec
	e.mux = http.NewServeMux()

	for _, rule := range e.spec.Rules {
		var hfunc http.HandlerFunc = nil
		e.mux.HandleFunc(rule.Accept, hfunc)
	}

	return
}

func (e *Default) GetHandler(ctx context.Context) (h http.Handler, err error) {
	h = e.mux
	return
}
