package models

import (
	"math"
	"sort"
)

type MenuSuggestionData struct {
	Session        Session
	SelectedItems  []MenuItem
	OrganisedItems []DetailedMenuItem
	RestaurantId   int64
}

type DetailedMenuItems []DetailedMenuItem

func (mi DetailedMenuItems) Len() int           { return len(mi) }
func (mi DetailedMenuItems) Less(i, j int) bool { return mi[i].Score < mi[j].Score }
func (mi DetailedMenuItems) Swap(i, j int)      { mi[i], mi[j] = mi[j], mi[i] }

func (suggestion *MenuSuggestionData) MakeMenuSuggestions() error {
	itemScores := make(map[int64]float64)
	tmpSelectedItems, err := suggestion.Session.GetSelectedMenuItems()
	if err != nil {
		return err
	}
	suggestion.SelectedItems = tmpSelectedItems
	allItems, err := suggestion.Session.GetDetailedSelectedMenuItemsForRestaurant(suggestion.Session.Id)
	if err != nil {
		return err
	}
	// Initialise and set distance modifier
	for _, item := range allItems {
		if _, ok := itemScores[item.Id]; ok == false {
			itemScores[item.Id] = 0
		}
	}
	for _, selection := range suggestion.SelectedItems {
		for _, item := range allItems {
			didMatch := false
			if selection.StyleOne == item.StyleOne || selection.StyleOne == item.StyleTwo || selection.StyleOne == item.StyleThree {
				itemScores[item.Id] += 1
				didMatch = true
			}
			if selection.StyleTwo == item.StyleOne || selection.StyleOne == item.StyleTwo || selection.StyleOne == item.StyleThree {
				itemScores[item.Id] += 1
				didMatch = true
			}
			if selection.StyleThree == item.StyleOne || selection.StyleOne == item.StyleTwo || selection.StyleOne == item.StyleThree {
				itemScores[item.Id] += 1
				didMatch = true
			}
			if selection.Flavor == item.Flavor {
				itemScores[item.Id] += 1
				didMatch = true
			}
			if didMatch {
				itemScores[item.Id] -= math.Abs(item.Heavy - selection.Heavy)
			}
		}
	}
	var menuItems DetailedMenuItems = make(DetailedMenuItems, len(itemScores))
	for ind, item := range allItems {
		menuItems[ind] = allItems[ind]
		menuItems[ind].Score = itemScores[item.Id]
	}
	sort.Sort(sort.Reverse(menuItems))
	suggestion.OrganisedItems = menuItems
	return nil
}
