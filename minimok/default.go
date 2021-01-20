package minimok

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/avrebarra/minimok/mokserver"
	"github.com/facebookgo/grace/gracehttp"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	MokSpecs []mokserver.Spec `validate:"required"`
}

type Minimok struct {
	config Config
}

func New(cfg Config) (*Minimok, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}
	return &Minimok{config: cfg}, nil
}

func (e *Minimok) Start(ctx context.Context) (err error) {
	// build handlers
	handlers, err := e.GetHandlers(ctx)
	if err != nil {
		return
	}

	// run servers and block
	for _, ms := range handlers {
		go func(ms Handler) {
			portaddr := ms.MokSpec.Port
			fmt.Printf("starting up mokserver '%s' on http://localhost:%d\n", ms.MokSpec.Name, portaddr)

			err = gracehttp.Serve(&http.Server{Addr: fmt.Sprint(":", portaddr), Handler: ms.Handler})
			if err != nil {
				err = fmt.Errorf("failed starting server: %s: %w", ms.MokSpec.Name, err)
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}(ms)
	}

	<-ctx.Done()
	fmt.Println("closing minimok...")

	return
}

func (e *Minimok) GetHandlers(ctx context.Context) (hs []Handler, err error) {
	hs = []Handler{}

	for _, sp := range e.config.MokSpecs {
		// build mokserver
		var m mokserver.MokServer
		if m, err = e.buildMokServer(sp); err != nil {
			return
		}

		// get handler
		var h http.Handler
		if h, err = m.GetHandler(ctx); err != nil {
			return
		}

		// assign
		hs = append(hs, Handler{
			MokSpec: sp,
			Handler: h,
		})
	}

	return
}

func (e *Minimok) buildMokServer(s mokserver.Spec) (m mokserver.MokServer, err error) {
	m = mokserver.New()
	err = m.ApplySpec(context.Background(), s)
	if err != nil {
		return
	}
	return
}
