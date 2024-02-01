package server

import (
	"context"
	"net/http"
	"wblzero/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *config.ServerHTTP, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		IdleTimeout:    cfg.IdleTimeout,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context)error{
return s.httpServer.Shutdown(ctx)
}