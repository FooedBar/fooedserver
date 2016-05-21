package models

type RestaurantSuggestion struct {
	SelectedItems        []MenuItem
	OrganisedRestaurants []Restaurant
}

func (suggestion *RestaurantSuggestion) MakeSuggestions() error {
	/*
		allRestaurants, err := GetAllRestaurants()
		if err != nil {
			return err
		}
		restaurantScores := map[int64]float64
		for _, r := range allRestaurants {
			restaurantScores[r.Id] = 0
		}
		allItems, err := GetAllMenuItems();
	*/
	return nil
}
