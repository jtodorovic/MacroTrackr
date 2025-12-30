package models

import (
	"time"

	"github.com/jtodorovic/macrotrackr/db"
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

// DTOs
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

// Mapper
func NewDailyGoalResponse(goal DailyGoal) DailyGoalResponse {
	return DailyGoalResponse{
		ID:             goal.ID,
		UserID:         goal.UserID,
		CaloriesTarget: goal.CaloriesTarget,
		ProteinTarget:  goal.ProteinTarget,
		CarbsTarget:    goal.CarbsTarget,
		FatsTarget:     goal.FatsTarget,
	}
}

func CreateDailyGoal(userID int64, req CreateDailyGoalRequest) (*DailyGoal, error) {
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

func GetLatestGoalForUser(userID int64) (*DailyGoal, error) {
	query := `SELECT * FROM daily_goals WHERE user_id = ? ORDER BY created_at DESC`
	row := db.DB.QueryRow(query, userID)

	var goal DailyGoal
	err := row.Scan(&goal.ID, &goal.UserID, &goal.CaloriesTarget, &goal.ProteinTarget, &goal.CarbsTarget, &goal.FatsTarget, &goal.CreatedAt, &goal.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &goal, nil
}

func GetDailyGoalByID(goalID int64) (*DailyGoal, error) {
	query := `SELECT * FROM daily_goals WHERE id = ?`
	row := db.DB.QueryRow(query, goalID)

	var goal DailyGoal
	err := row.Scan(&goal.ID, &goal.UserID, &goal.CaloriesTarget, &goal.ProteinTarget, &goal.CarbsTarget, &goal.FatsTarget, &goal.CreatedAt, &goal.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &goal, nil
}

// Apply update from UpdateDailyGoalRequest
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
