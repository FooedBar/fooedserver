package models

import (
	"fmt"
	"time"
)

type MenuItem struct {
	Id           int64     `json"id"`
	RestaurantId int64     `json:"restaurantId"`
	ImageUrl     string    `json:"imageUrl"`
	ImageHeight  int64     `json:"imageHeight"`
	ImageWidth   int64     `json:"imageWidth"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Flavor       string    `json:"flavor"`
	StyleOne     string    `json:"styleOne"`
	StyleTwo     string    `json:"styleTwo"`
	StyleThree   string    `json:"styleThree"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

var maxDistance float64 = 5.0 //Kilometres

type DetailedMenuItem struct {
	Id           int64  `json:"id"`
	RestaurantId int64  `json:"restaurantId"`
	ImageUrl     string `json:"imageUrl"`
	ImageHeight  string `json:"imageHeight"`
	ImageWidth   string `json:"imageWidth"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

func GetMenuItemsByPage(limit int, offset int, session Session) ([]DetailedMenuItem, error) {
	var items []DetailedMenuItem
	err := db.Raw("SELECT menu_item.id, menu_item.restaurant_id, menu_item.image_url, menu_item.image_height, menu_item.image_width, menu_item.name, menu_item.description, (6371 * acos( cos( radians(" + fmt.Sprintf("%.6f", session.CurrentLat) +
		") ) * cos( radians( restaurant.lat ) ) * cos( radians(restaurant.long) - radians(" + fmt.Sprintf("%.6f", session.CurrentLong) +
		")) + sin(radians(" + fmt.Sprintf("%.6f", session.CurrentLat) + "))" + " * sin( radians(restaurant.lat)))) AS distance FROM menu_item, restaurant WHERE menu_item.restaurant_id = restaurant.id HAVING distance < " + fmt.Sprintf("%.6f", maxDistance)).Scan(&items).Error
	return items, err
}
