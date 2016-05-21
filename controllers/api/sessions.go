package api

import (
	"encoding/json"
	"github.com/FooedBar/fooedserver/models"
	"net/http"
)

func V0_API_Create_Session(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqObj models.CreateSessionRequest
	err := decoder.Decode(&reqObj)
	if err != nil {
		failFormat := models.Response{Success: false, Debug: "Check JSON Formatting", Message: "Failed Creating the Session"}
		JSONResponse(w, failFormat, 400)
		return
	}
	if reqObj.Lat == 0 || reqObj.Long == 0 {
		failFormat := models.Response{Success: false, Debug: "Lat and long cannot be nil", Message: "Failed Creating the Session"}
		JSONResponse(w, failFormat, 400)
		return
	}
	session := models.Session{
		CurrentLat:  reqObj.Lat,
		CurrentLong: reqObj.Long,
	}
	err = session.Create()
	if err != nil {
		failParams := models.Response{Success: false, Debug: "Internal server error. If this issue persists, please submit a bug report to FooedBar", Message: "Failed Creating an Session"}
		JSONResponse(w, failParams, 500)
		return
	}
	successResp := models.Response{Success: true, Data: session, Message: "Successfully Created the Session"}
	JSONResponse(w, successResp, 200)
}
