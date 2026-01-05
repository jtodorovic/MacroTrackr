package services

import "github.com/jtodorovic/macrotrackr/internal/models"

func ToSummaryResponse(summary models.MacrosSummary) models.MacrosSummaryResponse {
	remainingCalories := summary.CaloriesTarget - summary.CaloriesConsumed
	remainingProtein := summary.ProteinTarget - summary.ProteinConsumed
	remainingCarbs := summary.CarbsTarget - summary.CarbsConsumed
	remainingFats := summary.FatsTarget - summary.FatsConsumed

	if remainingCalories < 0 {
		remainingCalories = 0
	}
	if remainingProtein < 0 {
		remainingProtein = 0
	}
	if remainingCarbs < 0 {
		remainingCarbs = 0
	}
	if remainingFats < 0 {
		remainingFats = 0
	}

	return models.MacrosSummaryResponse{
		UserID: summary.UserID,
		Date:   summary.Date,
		Goal: models.Macros{
			Calories: summary.CaloriesTarget,
			Protein:  summary.ProteinTarget,
			Carbs:    summary.CarbsTarget,
			Fats:     summary.FatsTarget,
		},
		Consumed: models.Macros{
			Calories: summary.CaloriesConsumed,
			Protein:  summary.ProteinConsumed,
			Carbs:    summary.CarbsConsumed,
			Fats:     summary.FatsConsumed,
		},
		Remaining: models.Macros{
			Calories: remainingCalories,
			Protein:  remainingProtein,
			Carbs:    remainingCarbs,
			Fats:     remainingFats,
		},
	}
}
