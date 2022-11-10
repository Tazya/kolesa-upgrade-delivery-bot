package main

import (
	"flag"
	"kolesa-upgrade-team/delivery-bot/internal/config"
	"kolesa-upgrade-team/delivery-bot/internal/delivery"
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("config: %s\n", err)
	}

	handler := delivery.NewHandler()
	server := server.NewServer(*cfg, handler)

	log.Printf("Starting server...\nhttp://localhost:%v\n", cfg.Addr)
	server.Srv.ListenAndServe()
}
