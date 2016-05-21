package models

import "time"

type Restaurant struct {
	Id          int64
	Lat         float64
	Long        float64
	Name        string
	Description string
	Style       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
