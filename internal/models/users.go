package models

import (
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	TelegramId int64  `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ChatId     int64  `json:"chat_id"`
}

type UserModel struct {
	Db *gorm.DB
}

func (m *UserModel) Create(user User) error {
	result := m.Db.Create(&user)

	return result.Error
}

func (u *User) Recipient() string {
	return strconv.Itoa(int(u.ChatId))
}

func (m *UserModel) GetAllUsers() ([]User, error) {
	var users []User

	res := m.Db.Find(users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}
