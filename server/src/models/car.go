package models

import "time"

type Car struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Brand          string    `json:"brand" binding:"required"`
	Model          string    `json:"model" binding:"required"`
	Year           uint      `json:"year" binding:"required,min=1886,max=2050"`
	DailyPrice     uint      `json:"daily_price" binding:"required"`
	InsurancePrice uint      `json:"insurance_price" binding:"required"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Rents          []Rent    `json:"rents" gorm:"foreignKey:CarID"`
}
