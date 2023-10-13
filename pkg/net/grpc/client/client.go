package client

import (
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/fantasy9830/go-graceful"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var connMap sync.Map

// Dial creates a client connection to the given target.
func Dial(target string) *grpc.ClientConn {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithMax(10), // 重試最大上限
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(3 * time.Second)), // 重試間隔時間
	}

	clientConn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		slog.Error("did not connect", "target", target, "err", err)
		os.Exit(1)
	}

	m := graceful.GetManager()
	m.RegisterOnShutdown(func() error {
		slog.Info("received an interrupt signal, close the gRPC client.")
		if err := clientConn.Close(); err != nil {
			slog.Error("failed to close gRPC client", "err", err)
			return err
		}

		return nil
	})

	return clientConn
}

// GetConn Initialized and exposed a singleton object of gRPC client.
func GetConn(target string) *grpc.ClientConn {
	if conn, ok := connMap.Load(target); ok {
		return conn.(*grpc.ClientConn)
	}

	conn := Dial(target)
	connMap.Store(target, conn)

	return conn
}
