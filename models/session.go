package models

import (
	"fmt"
	"time"
)

type Session struct {
	Id          int64     `json:"id"`
	CurrentLat  float64   `json:"-"`
	CurrentLong float64   `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func FindSessionById(id int64) (Session, error) {
	session := Session{}
	err := db.Where("id = ?", id).Limit(1).Find(&session).Error
	return session, err
}

func (session *Session) IsValid() bool {
	return session.Id > 0
}

func (session *Session) Create() error {
	return db.Create(&session).Error
}

func (session *Session) GetSelectedMenuItems() ([]MenuItem, error) {
	var selections []MenuItem
	err := db.Raw("SELECT menu_item.* FROM selection, menu_item WHERE selection.session_id = " + fmt.Sprint(session.Id) + " AND menu_item.id = selection.menu_item_id").Scan(&selections).Error
	return selections, err
}

func (session *Session) GetDetailedSelectedMenuItems() ([]DetailedMenuItem, error) {
	var selections []DetailedMenuItem
	err := db.Raw("SELECT menu_item.* FROM selection, menu_item WHERE selection.session_id = " + fmt.Sprint(session.Id) + " AND menu_item.id = selection.menu_item_id").Scan(&selections).Error
	return selections, err
}

func (session *Session) GetDetailedSelectedMenuItemsForRestaurant(restaurantId int64) ([]DetailedMenuItem, error) {
	var selections []DetailedMenuItem
	err := db.Raw("SELECT menu_item.*, (6371 * acos( cos( radians(" + fmt.Sprintf("%.6f", session.CurrentLat) +
		") ) * cos( radians( restaurant.lat ) ) * cos( radians(restaurant.long) - radians(" + fmt.Sprintf("%.6f", session.CurrentLong) +
		")) + sin(radians(" + fmt.Sprintf("%.6f", session.CurrentLat) + "))" + " * sin( radians(restaurant.lat)))) AS distance FROM menu_item, restaurant WHERE menu_item.restaurant_id = restaurant.id AND restaurant.id = " + fmt.Sprint(restaurantId)).Scan(&selections).Error
	return selections, err
}
