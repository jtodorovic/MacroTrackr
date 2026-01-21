package models

import (
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

type Macros struct {
	Calories int32 `json:"calories"`
	Protein  int32 `json:"protein"`
	Carbs    int32 `json:"carbs"`
	Fats     int32 `json:"fats"`
}

type MacrosSummaryResponse struct {
	UserID    int64     `json:"userId"`
	Date      time.Time `json:"date"`
	Goal      Macros    `json:"goal"`
	Consumed  Macros    `json:"consumed"`
	Remaining Macros    `json:"remaining"`
}
