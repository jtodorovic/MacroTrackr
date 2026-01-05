package services

import (
	"github.com/jtodorovic/macrotrackr/internal/db"
	"github.com/jtodorovic/macrotrackr/internal/models"
)

func NewDailyGoalResponse(goal models.DailyGoal) models.DailyGoalResponse {
	return models.DailyGoalResponse{
		ID:             goal.ID,
		UserID:         goal.UserID,
		CaloriesTarget: goal.CaloriesTarget,
		ProteinTarget:  goal.ProteinTarget,
		CarbsTarget:    goal.CarbsTarget,
		FatsTarget:     goal.FatsTarget,
	}
}

func CreateDailyGoal(userID int64, req models.CreateDailyGoalRequest) (*models.DailyGoal, error) {
	query := `
	INSERT INTO daily_goals(user_id, calories_target, protein_target, carbs_target, fats_target)
	VALUES (?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(userID, req.CaloriesTarget, req.ProteinTarget, req.CarbsTarget, req.FatsTarget)
	if err != nil {
		return nil, err
	}

	ID, err := result.LastInsertId()

	return GetDailyGoalByID(ID)
}

func GetLatestGoalForUser(userID int64) (*models.DailyGoal, error) {
	query := `SELECT * FROM daily_goals WHERE user_id = ? ORDER BY created_at DESC`
	row := db.DB.QueryRow(query, userID)

	var goal models.DailyGoal
	err := row.Scan(&goal.ID, &goal.UserID, &goal.CaloriesTarget, &goal.ProteinTarget, &goal.CarbsTarget, &goal.FatsTarget, &goal.CreatedAt, &goal.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &goal, nil
}

func GetDailyGoalByID(goalID int64) (*models.DailyGoal, error) {
	query := `SELECT * FROM daily_goals WHERE id = ?`
	row := db.DB.QueryRow(query, goalID)

	var goal models.DailyGoal
	err := row.Scan(&goal.ID, &goal.UserID, &goal.CaloriesTarget, &goal.ProteinTarget, &goal.CarbsTarget, &goal.FatsTarget, &goal.CreatedAt, &goal.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &goal, nil
}
