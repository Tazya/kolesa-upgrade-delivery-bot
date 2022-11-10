package server

import (
	"kolesa-upgrade-team/delivery-bot/internal/config"
	"kolesa-upgrade-team/delivery-bot/internal/delivery"
	"net/http"
	"time"
)

type Server struct {
	Srv *http.Server
}

func NewServer(cfg config.Config, handler *delivery.Handler) *Server {
	mux := http.NewServeMux()
	handler.InitRoutes(mux)

	return &Server{
		Srv: &http.Server{
			Addr:         ":" + cfg.Addr,
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}
