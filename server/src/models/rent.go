package models

import "time"

type Rent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ClientID  uint      `json:"client_id" binding:"required"`
	CarID     uint      `json:"car_id" binding:"required"`
}
