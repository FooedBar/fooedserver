package models

type Selection struct {
	Id         int64 `json:"id"`
	SessionId  int64 `json:"-"`
	MenuItemId int64 `json:"menu_item_id"`
	Like       bool  `json:"like"`
}

func (s *Selection) Create() error {
	return db.Create(&s).Error
}