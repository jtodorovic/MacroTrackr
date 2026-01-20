package services

import (
	"fmt"
	"math"

	"github.com/jtodorovic/macrotrackr/internal/integrations/usda"
	"github.com/jtodorovic/macrotrackr/internal/models"
)

const (
	NutrientCalories = 1008
	NutrientProtein  = 1003
	NutrientCarbs    = 1005
	NutrientFat      = 1004
)

func FetchNutrition(food string) (*usda.Food, error) {
	result, err := usda.SearchFood(food, 1)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	if len(result.Foods) == 0 {
		return nil, fmt.Errorf("No food found")
	}

	return &result.Foods[0], nil
}

func ExtractMacros(food *usda.Food, weight int32) models.Macros {
	var m models.Macros

	for _, n := range food.Nutrients {
		switch n.NutrientID {
		case NutrientCalories:
			m.Calories = int32(math.Round(n.Value))
		case NutrientProtein:
			m.Protein = int32(math.Round(n.Value))
		case NutrientCarbs:
			m.Carbs = int32(math.Round(n.Value))
		case NutrientFat:
			m.Fats = int32(math.Round(n.Value))
		}
	}

	// USDA values are per 100g
	factor := weight / 100.0

	return models.Macros{
		Calories: m.Calories * factor,
		Protein:  m.Protein * factor,
		Carbs:    m.Carbs * factor,
		Fats:     m.Fats * factor,
	}
}
