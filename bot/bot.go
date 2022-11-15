package bot

import (
	"kolesa-upgrade-team/delivery-bot/internal/models"
	"kolesa-upgrade-team/delivery-bot/usecase"
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

type ModifiedBot struct {
	Bot  *telebot.Bot
	user *models.UserModel
}

func NewModifiedBot(bot *telebot.Bot, u *models.UserModel) *ModifiedBot {
	return &ModifiedBot{
		Bot:  bot,
		user: u,
	}
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

func (bot *ModifiedBot) SendAll(msg usecase.Message) error {
	users, err := bot.user.GetAllUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		u := user
		_, err := bot.Bot.Send(&u, msg)
		if err != nil {
			return err
		}
	}

	return nil
}
