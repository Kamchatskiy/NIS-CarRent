package models

type Rent struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Price     uint   `json:"price"`
}
