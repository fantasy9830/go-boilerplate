package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/routes"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

// Server Server
type Server struct {
	*http.Server

	idleConnection chan struct{}
	eg             errgroup.Group
}

// NweServer NweServer
func NweServer() *Server {
	server := &Server{
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%s", config.Server.Port),
			Handler:      routes.Load(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		idleConnection: make(chan struct{}),
	}

	if config.Server.HTTPS {
		tlsConfig, err := TLSConfig()
		if err != nil {
			log.Fatal(err)
		}

		server.Server.TLSConfig = tlsConfig
	}

	return server
}

// TLSConfig TLSConfig
func TLSConfig() (*tls.Config, error) {
	if config.Server.Cert != "" && config.Server.Key != "" {
		cert, err := tls.LoadX509KeyPair(config.Server.Cert, config.Server.Key)
		if err != nil {
			return nil, err
		}

		return &tls.Config{
			Certificates: []tls.Certificate{cert},
		}, nil
	}

	parsed, err := url.Parse(config.Server.Host)
	if err != nil {
		return nil, err
	}

	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(parsed.Host),
		Cache:      autocert.DirCache("assets/key/server"),
	}

	return &tls.Config{
		GetCertificate: certManager.GetCertificate,
	}, nil
}

// Start Start Server
func (srv *Server) Start() error {
	go srv.graceful()

	srv.eg.Go(func() error {
		return srv.Server.ListenAndServe()
	})

	if err := srv.eg.Wait(); err != nil {
		return err
	}

	<-srv.idleConnection

	return nil
}

// StartTLS Start TLS Server
func (srv *Server) StartTLS() error {
	go srv.graceful()

	srv.eg.Go(func() error {
		return srv.Server.ListenAndServeTLS("", "")
	})

	if err := srv.eg.Wait(); err != nil {
		return err
	}

	<-srv.idleConnection

	return nil
}

func (srv *Server) graceful() {
	sigint := make(chan os.Signal, 1)

	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigint)
	<-sigint

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown: %v", err)
	}

	close(srv.idleConnection)
}
