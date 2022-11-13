package models

import "strconv"

func (u *User) Recipient() string {
	return strconv.Itoa(int(u.ChatId))
}