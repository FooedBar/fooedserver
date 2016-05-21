package api

import ()
import (
	"encoding/json"
	"github.com/FooedBar/fooedserver/models"
	"github.com/FooedBar/fooedserver/utils"
	"net/http"
)

func V0_API_Post_Selection(w http.ResponseWriter, r *http.Request) {
	var selectionObj models.Selection
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&selectionObj); err != nil {
		JSONResponse(w, models.Response{
			Success: false,
			Debug:   "Check JSON formatting",
			Message: "Failed to create selection",
		}, 400)
		return
	}
	session, err := utils.GetCurrentSession(r)
	if err != nil {
		JSONResponse(w, models.Response{
			Success: false,
			Debug:   "Internal Server Error. If this persists please submit a bug report to Fooedbar",
			Message: "Failed to create selection",
		}, 500)
		return
	}
	selectionObj.SessionId = session.Id
	if err := selectionObj.Create(); err != nil {
		JSONResponse(w, models.Response{
			Success: false,
			Debug:   "Internal Server Error. If this persists please submit a bug report to Fooedbar",
			Message: "Failed to create selection",
		}, 500)
		return
	}
	JSONResponse(w, models.Response{
		Success: true,
		Message: "Successfully created selection",
	}, 200)
}
