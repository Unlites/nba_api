package server

import (
	"net/http"
	"time"

	"github.com/Unlites/nba_api/config"
)

type Server struct {
	cfg    *config.Config
	router http.Handler
}

func NewServer(cfg *config.Config, router http.Handler) *Server {
	return &Server{cfg: cfg, router: router}
}

func (s *Server) Run() error {
	httpServer := &http.Server{
		Addr:           ":" + s.cfg.Server.Port,
		Handler:        s.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    s.cfg.Server.ReadTimeout * time.Second,
		WriteTimeout:   s.cfg.WriteTimeout * time.Second,
	}
	return httpServer.ListenAndServe()
}
