package http

import (
	"context"
	"net/http"
	"time"

	"warehouse/internal/config"
)

type Server struct {
	server http.Server
}

func New(cfg config.HTTPServer, handler http.Handler) *Server {

	return &Server{
		server: http.Server{
			Addr:                         cfg.Address,
			Handler:                      handler,
			DisableGeneralOptionsHandler: false,
			TLSConfig:                    nil,
			ReadTimeout:                  cfg.ReadTimeout,
			ReadHeaderTimeout:            0,
			WriteTimeout:                 0,
			IdleTimeout:                  cfg.IdleTimeout,
			MaxHeaderBytes:               0,
			TLSNextProto:                 nil,
			ConnState:                    nil,
			ErrorLog:                     nil,
			BaseContext:                  nil,
			ConnContext:                  nil,
		},
	}
}

func (s *Server) Run(ctx context.Context) {
	go s.server.ListenAndServe()
	<-ctx.Done()

	ctx1, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	s.server.Shutdown(ctx1)
}
