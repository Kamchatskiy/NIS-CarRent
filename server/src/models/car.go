package models

type Car struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	Year           uint   `json:"year"`
	DailyPrice     uint   `json:"daily_price"`
	InsurancePrice uint   `json:"insurance_price"`
}
