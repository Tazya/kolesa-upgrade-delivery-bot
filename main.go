package main

import (
	"flag"
	"kolesa-upgrade-team/delivery-bot/bot"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	modifiedBot := bot.ModifiedBot{
		Bot: bot.InitBot(cfg.BotToken),
	}

	modifiedBot.Bot.Handle("/hello", modifiedBot.HelloHandler)
	modifiedBot.Bot.Start()
}
