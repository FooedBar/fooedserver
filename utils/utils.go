package utils

import (
	"errors"
	"github.com/FooedBar/fooedserver/models"
	"github.com/gorilla/context"
	"net/http"
)

func SetCurrentSession(r *http.Request, installation models.Session) {
	context.Set(r, "session", installation)
}

func ClearCurrentRequest(r *http.Request) {
	context.Clear(r)
}

func GetCurrentSession(r *http.Request) (models.Session, error) {
	sessionObj := context.Get(r, "session")
	if sessionObj == nil {
		return models.Session{}, errors.New("Could not find the session, likely because of an invalid id.")
	}
	return sessionObj.(models.Session), nil
}
