package models

import "time"

type Task struct {
	ID          uint      `gorm:"id, primarykey" json:"-"`
	UserID      uint      `json:"user_id" gorm:"user_id"`
	Description string    `json:"description" gorm:"description"`
	Hours       int       `json:"-" gorm:"hours"`
	Minutes     int       `json:"-" gorm:"minutes"`
	StartedAt   time.Time `json:"-" gorm:"autoCreateTime"`
	Completed   bool      `json:"-" gorm:"completed"`
	CreatedAt   time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"-" gorm:"autoCreateTime"`
}
