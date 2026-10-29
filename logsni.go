package traefikplugin

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type Plugin struct {
	next http.Handler
	name string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Plugin{
		next: next,
		name: name,
	}, nil
}

func (a *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.TLS != nil {
		os.Stdout.WriteString(fmt.Sprintf("DANIEL-TRAEFIK-PLUGIN: tls_sni:%s", req.TLS.ServerName))
	} else {
		os.Stdout.WriteString(fmt.Sprintf("DANIEL-TRAEFIK-PLUGIN: request lacking TLS"))
	}
	a.next.ServeHTTP(rw, req)
}
