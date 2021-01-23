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
	handlers := []Handler{}
	for _, sp := range e.config.MokSpecs {
		h, ferr := buildHandler(sp)
		if ferr != nil {
			err = ferr
			return
		}

		// assign
		handlers = append(handlers, Handler{
			MokSpec: sp,
			Handler: h,
		})
	}

	// run servers and block
	for _, han := range handlers {
		go func(han Handler) {
			portaddr := han.MokSpec.Port
			fmt.Printf("starting up proxy server '%s' on http://localhost:%d\n", han.MokSpec.Name, portaddr)

			err = gracehttp.Serve(&http.Server{Addr: fmt.Sprint(":", portaddr), Handler: han.Handler})
			if err != nil {
				err = fmt.Errorf("failed starting server: %s: %w", han.MokSpec.Name, err)
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}(han)
	}

	<-ctx.Done()
	fmt.Println("closing minimok...")

	return
}

func buildHandler(sp mokserver.Spec) (h http.Handler, err error) {
	// build mokserver
	m := mokserver.New()
	err = m.ApplySpec(context.Background(), sp)
	if err != nil {
		return
	}

	// get handler
	if h, err = m.GetHandler(context.TODO()); err != nil {
		return
	}

	return
}
