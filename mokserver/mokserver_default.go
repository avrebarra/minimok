package mokserver

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/cssivision/reverseproxy"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Default struct {
	spec Spec
	mux  http.Handler
}

func New() MokServer {
	return &Default{
		mux:  http.DefaultServeMux,
		spec: Spec{},
	}
}

func (e *Default) ApplySpec(ctx context.Context, spec Spec) (err error) {
	e.spec = spec

	r := mux.NewRouter()

	for _, rule := range e.spec.Rules {
		var hfunc http.Handler = e.buildHandlerFunc(rule)

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

func (e *Default) buildHandlerFunc(rule MokSpecRule) (hf http.HandlerFunc) {
	hf = func(w http.ResponseWriter, r *http.Request) {
		// determine latency distribution
		latTotal := decideLatency(rule.MockLatency)
		var latEarly, latLate time.Duration

		switch rule.MockLatency.HogMode {
		case "e":
		case "early":
			latEarly = latTotal
			latLate = 0
			break

		case "l":
		case "late":
			latEarly = 0
			latLate = latTotal
			break

		default:
			latEarly = latTotal / 2
			latLate = latTotal / 2
		}

		// early latency
		time.Sleep(latEarly)
		if errors.Is(r.Context().Err(), context.Canceled) {
			return
		}

		// proxy if use origin specified
		if rule.UseOrigin != "" {
			// determine target path
			prefix := path.Clean(rule.Accept)
			r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)

			// parse target url
			path, err := url.Parse(rule.UseOrigin)
			if err != nil {
				panic(err)
			}

			// run proxy
			delayedwr := DelayedResponseWriter{ResponseWriter: w, Delay: latLate}
			proxy := reverseproxy.NewReverseProxy(path)
			proxy.ServeHTTP(delayedwr, r)

			return
		}

		// late latency
		time.Sleep(latLate)

		// return mock data
		w.WriteHeader(rule.MockResponse.Status)
		for k, v := range rule.MockResponse.Headers {
			w.Header().Set(k, v)
		}

		w.Write([]byte(rule.MockResponse.Body))

		return
	}

	return
}
