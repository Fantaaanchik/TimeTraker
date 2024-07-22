package models

import "time"

type User struct {
	ID             uint      `gorm:"id, primarykey" json:"-"`
	PassportNumber string    `json:"passport_number" gorm:"uniqueIndex"`
	Surname        string    `json:"surname" gorm:"surname"`
	Name           string    `json:"name" gorm:"name"`
	Patronymic     string    `json:"patronymic" gorm:"patronymic"`
	Address        string    `json:"address" gorm:"address"`
	CreatedAt      time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"-" gorm:"autoCreateTime"`
}

type CreateUser struct {
	PassportNumber string `json:"passport_number"`
}
