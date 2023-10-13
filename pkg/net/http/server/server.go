package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/fantasy9830/go-graceful"
)

// Server HTTP server object
type Server struct {
	*http.Server
}

// NewServer Create a new HTTP server object
func NewServer(handler http.Handler, options ...OptionFunc) *Server {
	srv := &Server{
		Server: &http.Server{
			Addr:    ":80",
			Handler: handler,
		},
	}

	for _, f := range options {
		f(srv)
	}

	return srv
}

// Start Start HTTP server
func (srv *Server) Start() {
	m := graceful.GetManager()

	m.Go(func(ctx context.Context) error {
		slog.Info("Starting HTTP server", "address", srv.Addr)
		if err := srv.Server.ListenAndServe(); err != nil {
			slog.Error(err.Error())
			return err
		}

		return nil
	})

	m.RegisterOnShutdown(func() error {
		slog.Info("received an interrupt signal, shut down the server.")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("HTTP server shutdown", "err", err)
			return err
		}

		return nil
	})
}
