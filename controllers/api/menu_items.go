package api

import (
	"encoding/json"
	"github.com/FooedBar/fooedserver/models"
	"github.com/FooedBar/fooedserver/utils"
	"net/http"
	"strconv"
)

func V0_API_Get_Menu_Items(w http.ResponseWriter, r *http.Request) {
	var reqObj models.GenericLimitOffsetRequest
	if r.Header.Get("Content-Type") == "application/json" {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&reqObj)
		if err != nil {
			failFormat := models.Response{Success: false, Debug: "Check JSON Formatting", Message: "Failed Getting the Menu Items"}
			JSONResponse(w, failFormat, 400)
			return
		}
	} else {
		reqObj.Offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
		reqObj.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}
	reqObj.Parameters()
	session, err := utils.GetCurrentSession(r)
	if err != nil {
		failParams := models.Response{Success: false, Debug: "Invalid Session Id", Message: "Failed Getting the Menu Items"}
		JSONResponse(w, failParams, 400)
		return
	}
	menuItems, err := models.GetMenuItemsByPage(reqObj.Limit, reqObj.Offset, session)
	if err != nil {
		failParams := models.Response{Success: false, Debug: "Internal server error. If this issue persists, please submit a bug report to FooedBar", Message: "Failed Getting the Menu Items"}
		JSONResponse(w, failParams, 500)
		return
	}
	inProgressResp := models.Response{Success: true, Data: menuItems, Message: "Successfully Retrieved the Menu Items"}
	JSONResponse(w, inProgressResp, 200)
}
