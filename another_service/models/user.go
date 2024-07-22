package models

import "time"

type UserInformation struct {
	ID             uint   `json:"id" gorm:"id, primarykey"`
	PassportNumber string `json:"passport_number" gorm:"passport_number"`
	Surname        string `json:"surname" gorm:"surname"`
	Name           string `json:"name" gorm:"name"`
	Patronymic     string `json:"patronymic" gorm:"patronymic"`
	Address        string `json:"address" gorm:"address"`
}

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
