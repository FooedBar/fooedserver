package api

import (
	"github.com/FooedBar/fooedserver/models"
	"github.com/gorilla/mux"
	"net/http"
)

func V0_API_Get_Restaurant(w http.ResponseWriter, r *http.Request) {
	restaurantId := mux.Vars(r)["id"]
	restaurant, err := models.GetRestaurantById(restaurantId)
	if err != nil {
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Failed to get restaurant",
			Debug:   "Internal server error. If this issue persists, please submit a bug report to FooedBar",
		}, 500)
		return
	}
	JSONResponse(w, models.Response{
		Success: true,
		Message: "Successfully retrieved restaurant",
		Data:    restaurant,
	}, 200)
}

/*
func V0_API_Get_Restaurants(w http.Response, r *http.Request) {
	_, err := utils.GetCurrentSession(r)
	if err != nil {
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Failed to get restaurants",
			Debug:   "Internal server error. If this issue persists, please submit a bug report to FooedBar",
		}, 500)
		return
	}
}
*/
