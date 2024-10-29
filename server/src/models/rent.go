package models

import "time"

type Rent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Price     uint      `json:"price" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ClientID  uint      `json:"client_id"`
	Client    Client    `json:"client" gorm:"foreignKey:ClientID"`
	CarID     uint      `json:"car_id"`
	Car       Car       `json:"car" gorm:"foreignKey:CarID"`
}
