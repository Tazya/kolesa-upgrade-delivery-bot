package bot

import (
	"kolesa-upgrade-team/delivery-bot/internal/config"
)

func LaunchBot(cfg *config.Config) {

	modifiedBot := ModifiedBot{
		Bot: InitBot(cfg.BotToken),
	}

	modifiedBot.Bot.Handle("/hello", modifiedBot.HelloHandler)
	modifiedBot.Bot.Start()
}
