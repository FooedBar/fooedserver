package models

type Selection struct {
	Id         int64 `json:"id"`
	SessionId  int64 `json:"-"`
	MenuItemId int64 `json:"-"`
	Like       bool  `json:"like"`
}
