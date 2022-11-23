package bot

import (
	"fmt"
	"kolesa-upgrade-team/delivery-bot/internal/models"
	"kolesa-upgrade-team/delivery-bot/usecase"
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

type ModifiedBot struct {
	Bot   *telebot.Bot
	users *models.UserModel
}

func NewModifiedBot(bot *telebot.Bot, u *models.UserModel) *ModifiedBot {
	return &ModifiedBot{
		Bot:   bot,
		users: u,
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

func (bot *ModifiedBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}

	existUser, err := bot.users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Ошибка получения пользователя %v", err)
	}

	if existUser == nil {
		err := bot.users.Create(newUser)

		if err != nil {
			return ctx.Send("Ошибка при создании пользователя...")
		}
		ctx.Send("Пользователь успешно создан!")
	}

	return ctx.Send("Привет, " + ctx.Sender().FirstName)
}

func (bot *ModifiedBot) SendAll(msg usecase.Message) error {
	const usersLimit = 100

	usersCount, err := bot.users.GetUsersCount()
	if err != nil {
		log.Printf("error while getting users count: %s", err)
		return err
	}

	if usersCount < usersLimit {
		usersCount = usersLimit
	}

	for i := 0; i <= usersCount-usersLimit; i += usersLimit {
		users, err := bot.users.GetUsersWithLimit(usersLimit, i)

		if err != nil {
			log.Printf("error while getting users: %s", err)
			return err
		}

		for _, user := range users {
			u := user
			_, err := bot.Bot.Send(&u, fmt.Sprintf("%s\n\n%s", msg.Title, msg.Body))
			if err != nil {
				log.Printf("error while sending message to users: %s", err.Error())
				return err
			}
		}
	}

	return nil
}
