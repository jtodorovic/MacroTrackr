package models

import "time"

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

func ToSummaryResponse(summary MacrosSummary) MacrosSummaryResponse {
	remainingCalories := summary.CaloriesTarget - summary.CaloriesConsumed
	remainingProtein := summary.ProteinTarget - summary.ProteinConsumed
	remainingCarbs := summary.CarbsTarget - summary.CarbsConsumed
	remainingFats := summary.FatsTarget - summary.FatsConsumed

	return MacrosSummaryResponse{
		UserID: summary.UserID,
		Date:   summary.Date,
		Goal: Macros{
			Calories: summary.CaloriesTarget,
			Protein:  summary.ProteinTarget,
			Carbs:    summary.CarbsTarget,
			Fats:     summary.FatsTarget,
		},
		Consumed: Macros{
			Calories: summary.CaloriesConsumed,
			Protein:  summary.ProteinConsumed,
			Carbs:    summary.CarbsConsumed,
			Fats:     summary.FatsConsumed,
		},
		Remaining: Macros{
			Calories: remainingCalories,
			Protein:  remainingProtein,
			Carbs:    remainingCarbs,
			Fats:     remainingFats,
		},
	}
}
