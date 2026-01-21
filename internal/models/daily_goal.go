package models

import (
	"time"

	"github.com/jtodorovic/macrotrackr/internal/db"
)

type DailyGoal struct {
	ID             int64
	UserID         int64
	CaloriesTarget int32
	ProteinTarget  int32
	CarbsTarget    int32
	FatsTarget     int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateDailyGoalRequest struct {
	CaloriesTarget int32 `json:"caloriesTarget" binding:"required"`
	ProteinTarget  int32 `json:"proteinTarget" binding:"required"`
	CarbsTarget    int32 `json:"carbsTarget" binding:"required"`
	FatsTarget     int32 `json:"fatsTarget" binding:"required"`
}

type UpdateDailyGoalRequest struct {
	CaloriesTarget *int32 `json:"caloriesTarget,omitempty"`
	ProteinTarget  *int32 `json:"proteinTarget,omitempty"`
	CarbsTarget    *int32 `json:"carbsTarget,omitempty"`
	FatsTarget     *int32 `json:"fatsTarget,omitempty"`
}

type DailyGoalResponse struct {
	ID             int64 `json:"id"`
	UserID         int64 `json:"userId"`
	CaloriesTarget int32 `json:"caloriesTarget"`
	ProteinTarget  int32 `json:"proteinTarget"`
	CarbsTarget    int32 `json:"carbsTarget"`
	FatsTarget     int32 `json:"fatsTarget"`
}

func (goal *DailyGoal) ApplyUpdate(req UpdateDailyGoalRequest) {
	if req.CaloriesTarget != nil {
		goal.CaloriesTarget = *req.CaloriesTarget
	}
	if req.ProteinTarget != nil {
		goal.ProteinTarget = *req.ProteinTarget
	}
	if req.CarbsTarget != nil {
		goal.CarbsTarget = *req.CarbsTarget
	}
	if req.FatsTarget != nil {
		goal.FatsTarget = *req.FatsTarget
	}
}

func (g *DailyGoal) Update() error {
	query := `
	 UPDATE daily_goals
	 SET calories_target = ?, protein_target = ?, carbs_target = ?, fats_target = ?
	 WHERE id = ?
	 `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(g.CaloriesTarget, g.ProteinTarget, g.CarbsTarget, g.FatsTarget, g.ID)
	return err
}

func (g *DailyGoal) Delete() error {
	query := "DELETE FROM daily_goals WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(g.ID)

	return err
}
