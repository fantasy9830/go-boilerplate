package server

import "google.golang.org/grpc"

type OptionFunc func(*Server)
type Register func(*grpc.Server)

func ServerAddr(addr string) OptionFunc {
	return func(srv *Server) {
		srv.Addr = addr
	}
}

func RegisterServer(registerServer Register) OptionFunc {
	return func(srv *Server) {
		srv.RegisterServer = registerServer
	}
}
