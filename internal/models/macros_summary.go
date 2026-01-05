package models

import (
	"math"
	"time"
)

type MacrosSummary struct {
	UserID int64
	Date   time.Time

	CaloriesConsumed int32
	ProteinConsumed  int32
	CarbsConsumed    int32
	FatsConsumed     int32

	CaloriesTarget int32
	ProteinTarget  int32
	CarbsTarget    int32
	FatsTarget     int32
}

type MacrosSummaryResponse struct {
	UserID    int64     `json:"userId"`
	Date      time.Time `json:"date"`
	Goal      Macros    `json:"goal"`
	Consumed  Macros    `json:"consumed"`
	Remaining Macros    `json:"remaining"`
}

type Macros struct {
	Calories int32 `json:"calories"`
	Protein  int32 `json:"protein"`
	Carbs    int32 `json:"carbs"`
	Fats     int32 `json:"fats"`
}

const (
	NutrientCalories = 1008
	NutrientProtein  = 1003
	NutrientCarbs    = 1005
	NutrientFat      = 1004
)

type USDAFood struct {
	Description   string         `json:"description"`
	FoodNutrients []USDANutrient `json:"foodNutrients"`
}

type USDANutrient struct {
	NutrientID int     `json:"nutrientId"`
	Value      float64 `json:"value"`
}

func ExtractMacros(food *USDAFood, weight int) Macros {
	var m Macros

	for _, n := range food.FoodNutrients {
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
	factor := float64(weight) / 100.0

	return Macros{
		Calories: m.Calories * int32(math.Round(factor)),
		Protein:  m.Protein * int32(math.Round(factor)),
		Carbs:    m.Carbs * int32(math.Round(factor)),
		Fats:     m.Fats * int32(math.Round(factor)),
	}
}
