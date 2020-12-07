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
	Name  string        `yaml:"name"`
	Port  int           `yaml:"port"`
	Rules []MuxSpecRule `yaml:"rules"`
}

type MuxSpecRule struct {
	Accept       string `yaml:"accept"`
	UseOrigin    string `yaml:"use_origin"`
	MockResponse struct {
		Status  int               `yaml:"status"`
		Body    string            `yaml:"body"`
		Headers map[string]string `yaml:"header"`
	} `yaml:"mock_response"`
	MockLatency MuxSpecRuleLatency `yaml:"mock_latency"`
}

type MuxSpecRuleLatency struct {
	Mode  string `yaml:"mode"`
	Value int    `yaml:"value"`
	Swing int    `yaml:"swing"`
}
