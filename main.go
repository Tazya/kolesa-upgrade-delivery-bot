package main

import (
	"flag"
	"kolesa-upgrade-team/delivery-bot/bot"
	"kolesa-upgrade-team/delivery-bot/internal/config"
	"kolesa-upgrade-team/delivery-bot/internal/delivery"
	"kolesa-upgrade-team/delivery-bot/internal/models"
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	var wg sync.WaitGroup

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("config: %s\n", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database initialization error: %s\n", err)
	}

	b := bot.NewModifiedBot(bot.InitBot(cfg.BotToken), &models.UserModel{Db: db})
	handler := delivery.NewHandler(b)
	server := server.NewServer(*cfg, handler)

	log.Printf("Starting server...\nhttp://localhost:%v\n", cfg.Addr)

	wg.Add(2)
	go func() {
		server.Srv.ListenAndServe()
		wg.Done()
	}()
	go func() {
		bot.LaunchBot(b)
		wg.Done()
	}()
	wg.Wait()
}
