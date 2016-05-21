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

func GetRestaurantById(id string) (Restaurant, error) {
	var restaurant Restaurant
	err := db.Where("id = ?", id).First(&restaurant).Error
	return restaurant, err
}
