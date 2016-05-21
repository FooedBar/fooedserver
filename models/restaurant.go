package models

import "time"

type Restaurant struct {
	Id          int64     `json:"id"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"imageUrl"`
	ImageHeight int64     `json:"imageHeight"`
	ImageWidth  int64     `json:"imageWidth"`
	Description string    `json:"description"`
	Style       string    `json:"style"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func GetRestaurantById(id string) (Restaurant, error) {
	var restaurant Restaurant
	err := db.Where("id = ?", id).First(&restaurant).Error
	return restaurant, err
}
