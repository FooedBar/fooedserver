package models

import "time"

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
