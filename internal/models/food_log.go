package models

import (
	"time"

	"github.com/jtodorovic/macrotrackr/internal/db"
)

type FoodLog struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Name      string    `json:"name"`
	Calories  int32     `json:"calories"`
	Protein   int32     `json:"protein"`
	Carbs     int32     `json:"carbs"`
	Fats      int32     `json:"fats"`
	MealType  string    `json:"mealType"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateFoodLogRequest struct {
	Name     string `json:"name" binding:"required"`
	Calories int32  `json:"calories" binding:"required"`
	Protein  int32  `json:"protein" binding:"required"`
	Carbs    int32  `json:"carbs" binding:"required"`
	Fats     int32  `json:"fats" binding:"required"`
	MealType string `json:"mealType" binding:"required"`
}

type CreateFoodLogByWeightRequest struct {
	Food     string `json:"food" binding:"required"`
	Weight   int32  `json:"weight" binding:"required"`
	MealType string `json:"mealType" binding:"required"`
}

type UpdateFoodLogRequest struct {
	Name     *string `json:"name"`
	Calories *int32  `json:"calories"`
	Protein  *int32  `json:"protein"`
	Carbs    *int32  `json:"carbs"`
	Fats     *int32  `json:"fats"`
	MealType *string `json:"mealType"`
}

func (f *FoodLog) ApplyUpdate(req UpdateFoodLogRequest) {
	if req.Name != nil {
		f.Name = *req.Name
	}
	if req.Calories != nil {
		f.Calories = *req.Calories
	}
	if req.Protein != nil {
		f.Protein = *req.Protein
	}
	if req.Carbs != nil {
		f.Carbs = *req.Carbs
	}
	if req.Fats != nil {
		f.Fats = *req.Fats
	}
	if req.MealType != nil {
		f.MealType = *req.MealType
	}
}

func (f *FoodLog) Update() error {
	query := `
		UPDATE food_logs
		SET name = ?, calories = ?, protein = ?, carbs = ?, fats = ?, meal_type = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(
		query,
		f.Name,
		f.Calories,
		f.Protein,
		f.Carbs,
		f.Fats,
		f.MealType,
		f.ID,
	)

	return err
}

func (f *FoodLog) Delete() error {
	_, err := db.DB.Exec(`DELETE FROM food_logs WHERE id = ?`, f.ID)
	return err
}
