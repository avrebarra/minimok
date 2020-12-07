package minimok

import (
	"context"
	"net/http"

	"github.com/avrebarra/minimok/mux"
)

type Minimok interface {
	GetHandlers(ctx context.Context) (hs []MuxHandler, err error)
}

type MuxHandler struct {
	MuxSpec mux.MuxSpec
	Handler http.Handler
}
