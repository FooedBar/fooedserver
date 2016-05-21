package models

import "time"

type Restaurant struct {
	Id          int64
	Lat         float64
	Long        float64
	ImageUrl    string
	ImageHeight int64
	ImageWidth  int64
	Name        string
	Description string
	Style       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
