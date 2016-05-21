package models

import "time"

type MenuItem struct {
	Id           int64
	RestaurantId int64
	ImageUrl     string
	ImageHeight  int64
	ImageWidth   int64
	Flavor       string
	StyleOne     string
	StyleTwo     string
	StyleThree   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
