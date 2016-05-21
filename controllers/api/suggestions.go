package api

import (
	"github.com/FooedBar/fooedserver/models"
	"github.com/FooedBar/fooedserver/utils"
	"net/http"
)

func V0_API_Get_Restaurant_Suggestions(w http.ResponseWriter, r *http.Request) {
	session, err := utils.GetCurrentSession(r)
	if err != nil {
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Failed to get restaurants",
			Debug:   "Internal server error. If this issue persists, please submit a bug report to FooedBar",
		}, 500)
		return
	}
	request := models.RestaurantSuggestionData{
		Session: session,
	}
	err = request.MakeRestaurantSuggestions()
	if err != nil {
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Failed to get restaurants",
			Debug:   "Internal server error. If this issue persists, please submit a bug report to FooedBar",
		}, 500)
		return
	}
	wrapper := models.GenericItemArrayWrapper{
		Items: request.OrganisedRestaurants,
	}
	JSONResponse(w, models.Response{
		Success: true,
		Message: "Successfully retrieved restaurant",
		Data:    wrapper,
	}, 200)
}
