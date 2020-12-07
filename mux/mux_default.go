package mux

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Default struct {
	spec MuxSpec
	mux  http.Handler
}

func New() Mux {
	return &Default{
		mux:  http.DefaultServeMux,
		spec: MuxSpec{},
	}
}

func (e *Default) ApplySpec(ctx context.Context, spec MuxSpec) (err error) {
	e.spec = spec

	r := mux.NewRouter()

	for _, rule := range e.spec.Rules {
		var hfunc http.Handler = buildMuxSpecRuleHandlerFunc(rule)

		hfunc = handlers.CombinedLoggingHandler(os.Stdout, hfunc)
		hfunc = handlers.CombinedLoggingHandler(os.Stdout, hfunc)

		r.HandleFunc(rule.Accept, hfunc.ServeHTTP)
	}

	e.mux = r

	return
}

func (e *Default) GetHandler(ctx context.Context) (h http.Handler, err error) {
	var hfunc http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
		e.mux.ServeHTTP(rw, r)
	}

	h = hfunc

	return
}
