package car

import "time"

type Make struct {
	ID          int64
	Name        string
	CountryCode *string
	LogoURL     *string
}

type Model struct {
	ID                  int64
	MakeID              int64
	Name                string
	ProductionStartYear *int16
	ProductionEndYear   *int16
}

type Car struct {
	ID             int64     `json:"id"`
	Make           string    `json:"make"`
	Model          string    `json:"model"`
	VIN            *string   `json:"vin"`
	Year           int16     `json:"year"`
	CurrentMileage int32     `json:"current_mileage"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateCarParams struct {
	ModelID        int64
	VIN            *string
	Year           int16
	CurrentMileage int32
}

type CreateParams struct {
	Make           string
	Model          string
	VIN            *string
	Year           int16
	CurrentMileage int32
}
