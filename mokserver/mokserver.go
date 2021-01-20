package mokserver

import (
	"context"
	"net/http"
)

type MokServer interface {
	ApplySpec(ctx context.Context, spec Spec) (err error)
	GetHandler(ctx context.Context) (h http.Handler, err error)
}

type Spec struct {
	Name  string        `yaml:"name"`
	Port  int           `yaml:"port"`
	Rules []MokSpecRule `yaml:"rules"`
}

type MokSpecRule struct {
	Accept       string `yaml:"accept"`
	UseOrigin    string `yaml:"use_origin"`
	MockResponse struct {
		Status  int               `yaml:"status"`
		Body    string            `yaml:"body"`
		Headers map[string]string `yaml:"header"`
	} `yaml:"mock_response"`
	MockLatency MokSpecRuleLatency `yaml:"mock_latency"`
}

type MokSpecRuleLatency struct {
	Mode    string `yaml:"mode"`
	HogMode string `yaml:"hog"`
	Value   int    `yaml:"value"`
	Swing   int    `yaml:"swing"`
}
