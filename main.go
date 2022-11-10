package main

import (
	"flag"
	"kolesa-upgrade-team/delivery-bot/bot"
	"kolesa-upgrade-team/delivery-bot/internal/config"
	"kolesa-upgrade-team/delivery-bot/internal/delivery"
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
	"sync"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	var wg sync.WaitGroup
	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("config: %s\n", err)
	}

	handler := delivery.NewHandler()
	server := server.NewServer(*cfg, handler)

	log.Printf("Starting server...\nhttp://localhost:%v\n", cfg.Addr)
	wg.Add(2)
	go func() {
		server.Srv.ListenAndServe()
		wg.Done()
	}()
	go func() {
		bot.LaunchBot(cfg)
		wg.Done()
	}()
	wg.Wait()
}
