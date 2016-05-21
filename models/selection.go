package models

type Selection struct {
	Id         int64 `json:"-"`
	SessionId  int64 `json:"-"`
	MenuItemId int64 `json:"menuItemId"`
	Like       bool  `json:"like"`
}

func (s *Selection) Create() error {
	return db.Create(&s).Error
}
