package models

import (
	"time"

	"github.com/jtodorovic/macrotrackr/db"
)

type Goal struct {
	ID             int64     `json:"id"`
	UserID         string    `json:"userId" binding:"required"`
	CaloriesTarget int32     `json:"caloriesTarget" binding:"required"`
	ProteinTarget  int32     `json:"proteinTarget" binding:"required"`
	CarbsTarget    int32     `json:"carbsTarget" binding:"required"`
	FatsTarget     int32     `json:"fatsTarget" binding:"required"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (g Goal) Save() error {
	query := `
	INSERT INTO goals(user_id, calories_target, protein_target, carbs_target, fats_target)
	VALUES (?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	g.ID = id
	return err
}

var goals = []Goal{}

func GetGoalsForUser(userId int64) ([]Goal, error) {
	query := `SELECT * FROM goals WHERE user_id = ?`

	rows, err := db.DB.Query(query, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var goal Goal
		err := rows.Scan(&goal.ID, &goal.UserID, &goal.CaloriesTarget, &goal.ProteinTarget, &goal.CarbsTarget, &goal.FatsTarget)
		if err != nil {
			return nil, err
		}

		goals = append(goals, goal)
	}

	return goals, nil
}
