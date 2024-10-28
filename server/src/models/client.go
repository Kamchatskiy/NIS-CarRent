package models

import "time"

type Client struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required"`
	Surname     string    `json:"surname" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
