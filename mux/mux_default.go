package mux

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct{}

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
		fmt.Printf("registering: %+v\n", rule)

		var hfunc http.Handler = buildMuxSpecRuleHandlerFunc(rule)

		hfunc = handlers.CombinedLoggingHandler(os.Stdout, hfunc)
		hfunc = handlers.CombinedLoggingHandler(os.Stdout, hfunc)

		e.mux.Handle(rule.Accept, hfunc)
	}

	return
}

func (e *Default) GetHandler(ctx context.Context) (h http.Handler, err error) {
	var hfunc http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
		e.mux.ServeHTTP(rw, r)
	}

	h = hfunc

	return
}
