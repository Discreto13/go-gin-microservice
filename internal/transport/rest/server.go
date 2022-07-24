package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	readTimeout, writeTimeout = 10 * time.Second, 10 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port uint, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      handler,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
