package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique;" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"` //json:"-": this ensure that the user's password is not returned in the JSON response
	Entries  []EntryModel
}
