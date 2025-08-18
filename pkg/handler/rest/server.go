package rest

import (
	"context"
	"net/http"

	"github.com/hossein1376/gotp/pkg/service"
)

type Server struct {
	srv      *http.Server
	services *service.Services
}

func NewServer(addr string, srvc *service.Services) *Server {
	mux := newRouter(srvc)

	return &Server{
		services: srvc,
		srv: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
