package models

type Selection struct {
	Id         int64
	SessionId  int64
	MenuItemId int64
	Like       bool
}
