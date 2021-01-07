package mux

import (
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/cssivision/reverseproxy"
	"gopkg.in/yaml.v3"
)

func MuxSpecFromYAML(bits []byte) (sp MuxSpec, err error) {
	err = yaml.Unmarshal(bits, &sp)
	if err != nil {
		return
	}

	return
}

func buildMuxSpecRuleHandlerFunc(rule MuxSpecRule) (hf http.HandlerFunc) {
	hf = func(w http.ResponseWriter, r *http.Request) {
		// determine latency distribution
		latTotal := decideLatency(rule.MockLatency)
		var latEarly, latLate time.Duration

		switch rule.MockLatency.HogMode {
		case "early":
			latEarly = latTotal
			latLate = 0
			break

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
			proxy := reverseproxy.NewReverseProxy(path)
			proxy.ServeHTTP(w, r)

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
