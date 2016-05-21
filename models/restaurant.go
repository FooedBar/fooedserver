package models

import "time"

type Restaurant struct {
<<<<<<< HEAD
	Id          int64     `json:"id"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Style       string    `json:"style"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
=======
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
>>>>>>> 2ee5f97127ecbce38da3eb34b43b6206b0ab35d0
}

func GetRestaurantById(id string) (Restaurant, error) {
	var restaurant Restaurant
	err := db.Where("id = ?", id).First(&restaurant).Error
	return restaurant, err
}
