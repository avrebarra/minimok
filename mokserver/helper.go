package mokserver

import (
	"math/rand"
	"net/http"
	"time"
)

type DelayedResponseWriter struct {
	http.ResponseWriter
	Delay time.Duration
}

func (w DelayedResponseWriter) Write(b []byte) (int, error) {
	time.Sleep(w.Delay)
	return w.ResponseWriter.Write(b)
}

func decideLatency(e MokSpecRuleLatency) (d time.Duration) {
	switch e.Mode {
	case "c":
	case "const":
	case "constant":
		d = time.Duration(e.Value) * time.Millisecond
		break

	case "m":
	case "max":
	case "maximum":
		d = time.Duration(rand.Float64()*float64(e.Value)) * time.Millisecond
		break

	case "s":
	case "swing":
		d = time.Duration(rand.Float64()*float64(e.Swing)+float64(e.Value-e.Swing/2)) * time.Millisecond
		break
	}

	return
}
