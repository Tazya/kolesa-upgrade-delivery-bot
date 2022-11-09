package bot

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

type ModifiedBot struct {
	Bot *telebot.Bot
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}

func (bot *ModifiedBot) HelloHandler(ctx telebot.Context) error {
	return ctx.Send("Hello " + ctx.Sender().FirstName)
}
