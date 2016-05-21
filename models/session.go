package models

import "time"

type Session struct {
	Id          int64
	CurrentLat  float64
	CurrentLong float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FindSessionById(id int64) (Session, error) {
	session := Session{}
	err := db.Where("id = ?", id).Limit(1).Find(&session).Error
	return session, err
}

func (session *Session) IsValid() bool {
	return session.Id > 0
}
