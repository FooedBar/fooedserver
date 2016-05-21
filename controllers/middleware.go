package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/FooedBar/fooedserver/models"
	"github.com/FooedBar/fooedserver/utils"
	"net/http"
	"strconv"
)

// GetContext wraps each request in a function which fills in the context for a given request.
func GetContext(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.Header.Get("X-SESSION-ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			handler.ServeHTTP(w, r)
			utils.ClearCurrentRequest(r)
			return
		}
		if id > 0 {
			session, err := models.FindSessionById(int64(id))
			if err != nil {
				fmt.Println("FIRST CASE")
				JSONError(w, 400, "Missing X-SESSION-ID header")
				return
			} else {
				utils.SetCurrentSession(r, session)
				handler.ServeHTTP(w, r)
				utils.ClearCurrentRequest(r)
				return
			}
		} else {
			handler.ServeHTTP(w, r)
			utils.ClearCurrentRequest(r)
		}
	}
}

func RequireSessionId(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := utils.GetCurrentSession(r)
		if err != nil {
			fmt.Println("SECOND CASE")
			JSONError(w, 400, "Missing X-SESSION-ID header")
			return
		}
		if session.IsValid() == false {
			JSONError(w, 401, "Invalid X-SESSION-ID header value")
			return
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}

func JSONError(w http.ResponseWriter, c int, m string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	cj, _ := json.MarshalIndent(models.Response{Success: false, Message: m}, "", "  ")
	fmt.Fprintf(w, "%s", cj)
}
