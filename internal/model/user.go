package model

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	Username  string    `gorm:"type:varchar(255);unique"`
	Password  string    `gorm:"type:varchar(255);"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (User) TableName() string { return "user" }
