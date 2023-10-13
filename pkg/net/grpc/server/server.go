package server

import (
	"context"
	"errors"
	"log/slog"
	"net"

	"github.com/fantasy9830/go-graceful"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	grpc           *grpc.Server
	Addr           string
	RegisterServer Register
}

func NewServer(options ...OptionFunc) *Server {
	// Define customfunc to handle panic
	var customFunc grpc_recovery.RecoveryHandlerFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(opts...),
		),
	)

	srv := &Server{
		grpc: grpcServer,
		Addr: ":8080",
	}

	for _, f := range options {
		f(srv)
	}

	return srv
}

func (srv *Server) listenAndServe() error {
	lis, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}

	if srv.RegisterServer == nil {
		return errors.New("RegisterServer is required")
	}

	srv.RegisterServer(srv.grpc)

	return srv.grpc.Serve(lis)
}

// Start Start HTTP server
func (srv *Server) Start() {
	m := graceful.GetManager()

	m.Go(func(ctx context.Context) error {
		slog.Info("Starting gRPC server", "address", srv.Addr)
		if err := srv.listenAndServe(); err != nil {
			slog.Error(err.Error())
			return err
		}

		return nil
	})

	m.RegisterOnShutdown(func() error {
		slog.Info("received an interrupt signal, shut down gRPC server")
		srv.grpc.GracefulStop()

		return nil
	})
}
