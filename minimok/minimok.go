package minimok

import (
	"net/http"

	"github.com/avrebarra/minimok/mokserver"
)

type Handler struct {
	MokSpec mokserver.Spec
	Handler http.Handler
}
