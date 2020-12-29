package minimok

import (
	"context"
	"net/http"

	"github.com/avrebarra/minimok/mux"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	MuxSpecs []mux.MuxSpec `validate:"required"`
}

type MinimokStruct struct {
	config Config
}

func New(cfg Config) (Minimok, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}
	return &MinimokStruct{config: cfg}, nil
}

func (e *MinimokStruct) GetHandlers(ctx context.Context) (hs []MuxHandler, err error) {
	hs = []MuxHandler{}

	for _, sp := range e.config.MuxSpecs {
		// build mux
		var m mux.Mux
		if m, err = e.buildMux(sp); err != nil {
			return
		}

		// get handler
		var h http.Handler
		if h, err = m.GetHandler(ctx); err != nil {
			return
		}

		// assign
		hs = append(hs, MuxHandler{
			MuxSpec: sp,
			Handler: h,
		})
	}

	return
}

func (e *MinimokStruct) buildMux(s mux.MuxSpec) (m mux.Mux, err error) {
	m = mux.New()
	err = m.ApplySpec(context.Background(), s)
	if err != nil {
		return
	}
	return
}
