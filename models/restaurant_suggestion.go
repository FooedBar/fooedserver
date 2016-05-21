package models

import (
	"math"
	"sort"
)

type RestaurantSuggestionData struct {
	Session              Session
	SelectedItems        []MenuItem
	OrganisedRestaurants []Restaurant
}

type Restaurants []Restaurant

func (r Restaurants) Len() int           { return len(r) }
func (r Restaurants) Less(i, j int) bool { return r[i].Score < r[j].Score }
func (r Restaurants) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (suggestion *RestaurantSuggestionData) MakeRestaurantSuggestions() error {
	restaurantScores := make(map[int64]float64)
	restaurantDistances := make(map[int64]float64)
	tmpSelectedItems, err := suggestion.Session.GetSelectedMenuItems()
	if err != nil {
		return err
	}
	suggestion.SelectedItems = tmpSelectedItems
	allItems, err := GetAllMenuItemsWithinSession(suggestion.Session)
	if err != nil {
		return err
	}
	// Initialise and set distance modifier
	for _, item := range allItems {
		if _, ok := restaurantScores[item.RestaurantId]; ok == false {
			restaurantScores[item.RestaurantId] -= item.Distance * 3
			restaurantDistances[item.RestaurantId] = item.Distance
		}
	}
	for _, selection := range suggestion.SelectedItems {
		for _, item := range allItems {
			didMatch := false
			if selection.StyleOne == item.StyleOne || selection.StyleOne == item.StyleTwo || selection.StyleOne == item.StyleThree {
				restaurantScores[item.RestaurantId] += 1
				didMatch = true
			}
			if selection.StyleTwo == item.StyleOne || selection.StyleOne == item.StyleTwo || selection.StyleOne == item.StyleThree {
				restaurantScores[item.RestaurantId] += 1
				didMatch = true
			}
			if selection.StyleThree == item.StyleOne || selection.StyleOne == item.StyleTwo || selection.StyleOne == item.StyleThree {
				restaurantScores[item.RestaurantId] += 1
				didMatch = true
			}
			if selection.Flavor == item.Flavor {
				restaurantScores[item.RestaurantId] += 1
				didMatch = true
			}
			if didMatch {
				restaurantScores[item.RestaurantId] -= math.Abs(item.Heavy - selection.Heavy)
			}
		}
	}
	restaurantIds := make([]int64, len(restaurantScores))
	tmpCnt := 0
	for rId, _ := range restaurantScores {
		restaurantIds[tmpCnt] = rId
		tmpCnt++
	}
	var restaurants Restaurants
	restaurants, err = GetRestaurants(restaurantIds)
	if err != nil {
		return err
	}
	for ind, r := range restaurants {
		restaurants[ind].Score = restaurantScores[r.Id]
		restaurants[ind].Distance = restaurantDistances[r.Id]
	}
	sort.Sort(restaurants)
	suggestion.OrganisedRestaurants = restaurants
	return nil
}
