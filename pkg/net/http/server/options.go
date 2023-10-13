package server

type OptionFunc func(*Server)

func ServerAddr(addr string) OptionFunc {
	return func(srv *Server) {
		srv.Addr = addr
	}
}
