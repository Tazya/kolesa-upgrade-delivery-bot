package bot

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func LaunchBot(configPath *string) {

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	modifiedBot := ModifiedBot{
		Bot: InitBot(cfg.BotToken),
	}

	modifiedBot.Bot.Handle("/hello", modifiedBot.HelloHandler)
	modifiedBot.Bot.Start()
}
