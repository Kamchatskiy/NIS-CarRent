package models

import "time"

type Rent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ClientID  uint      `json:"client_id"`
	CarID     uint      `json:"car_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Price     uint      `json:"price"`
}
