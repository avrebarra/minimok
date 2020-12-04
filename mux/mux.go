package mux

import (
	"context"
	"net/http"
)

type Mux interface {
	ApplySpec(ctx context.Context, spec MuxSpec) (err error)
	GetHandler(ctx context.Context) (h http.Handler, err error)
}

type MuxSpec struct {
	Rules []MuxSpecRule `json:"rules"`
}

type MuxSpecRule struct {
	Accept       string `json:"accept"`
	UseOrigin    string `json:"use_origin"`
	MockResponse struct {
		Status  int               `json:"status"`
		Body    string            `json:"body"`
		Headers map[string]string `json:"header"`
	} `json:"mock_response"`
	MockLatency MuxSpecRuleLatency `json:"mock_latency"`
}

type MuxSpecRuleLatency struct {
	Mode  string `json:"mode"`
	Value int    `json:"value"`
	Swing int    `json:"swing"`
}
