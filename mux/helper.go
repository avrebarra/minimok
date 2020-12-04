package mux

import (
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/cssivision/reverseproxy"
)

func buildMuxSpecRuleHandlerFunc(rule MuxSpecRule) (hf http.HandlerFunc) {
	hf = func(w http.ResponseWriter, r *http.Request) {
		// adjust latencies
		time.Sleep(decideLatency(rule.MockLatency))

		// proxy if use origin specified
		if rule.UseOrigin != "" {
			path, err := url.Parse(rule.UseOrigin)
			if err != nil {
				panic(err)
			}
			proxy := reverseproxy.NewReverseProxy(path)
			proxy.ServeHTTP(w, r)
			return
		}

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

func decideLatency(e MuxSpecRuleLatency) (d time.Duration) {
	switch e.Mode {
	case "const":
		d = time.Duration(e.Value) * time.Millisecond
		break

	case "max":
		d = time.Duration(rand.Float64()*float64(e.Value)) * time.Millisecond
		break

	case "swing":
		d = time.Duration(rand.Float64()*float64(e.Swing)+float64(e.Value-e.Swing/2)) * time.Millisecond
		break
	}

	return
}
