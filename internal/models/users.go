package models

import (
	"strconv"

	"gorm.io/gorm"
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

func (m *UserModel) GetUsersWithLimit(limit int, offset int) ([]User, error) {
	var users []User

	res := m.Db.Limit(limit).Offset(offset).Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (m *UserModel) FindOne(telegramId int64) (*User, error) {
	existUser := User{}
	result := m.Db.First(&existUser, User{TelegramId: telegramId})

	if result.Error != nil {
		return nil, result.Error
	}

	return &existUser, nil
}

func (m *UserModel) GetUsersCount() (int, error) {
	var count int64

	res := m.Db.Table("users").Count(&count)
	if res.Error != nil {
		return 0, nil
	}

	return int(count), nil
}
