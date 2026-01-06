package services

import (
	"time"

	"github.com/jtodorovic/macrotrackr/internal/db"
	"github.com/jtodorovic/macrotrackr/internal/models"
)

func CreateFoodLog(userID int64, req models.CreateFoodLogByWeightRequest) (*models.FoodLog, error) {
	food, err := FetchNutrition(req.Food)
	if err != nil {
		return nil, err
	}

	macros := ExtractMacros(food, req.Weight)

	query := `
	INSERT INTO food_logs(user_id, name, calories, protein, carbs, fats, meal_type)
	VALUES (?,?,?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(userID, req.Food, macros.Calories, macros.Protein, macros.Carbs, macros.Fats, req.MealType)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindFoodLogByID(id)
}

func FindFoodLogByID(ID int64) (*models.FoodLog, error) {
	query := `SELECT id, user_id, name, calories, protein, carbs, fats, meal_type, created_at, updated_at
	          FROM food_logs 
			  WHERE id = ?`
	row := db.DB.QueryRow(query, ID)

	var foodLog models.FoodLog
	err := row.Scan(&foodLog.ID,
		&foodLog.UserID,
		&foodLog.Name,
		&foodLog.Calories,
		&foodLog.Protein,
		&foodLog.Carbs,
		&foodLog.Fats,
		&foodLog.MealType,
		&foodLog.CreatedAt,
		&foodLog.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &foodLog, nil
}

func startAndEndOfDay(date time.Time) (time.Time, time.Time) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := start.Add(24 * time.Hour)
	return start, end
}

func FindFoodLogsByUserAndDate(userID int64, date time.Time) ([]models.FoodLog, error) {
	startDate, endDate := startAndEndOfDay(date)
	query := `SELECT id, user_id, name, calories, protein, carbs, fats, meal_type, created_at
        FROM food_logs
        WHERE user_id = ?
        AND created_at >= ?
        AND created_at < ?
        ORDER BY created_at ASC
    `
	rows, err := db.DB.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.FoodLog

	for rows.Next() {
		var log models.FoodLog
		err := rows.Scan(
			&log.ID,
			&log.UserID,
			&log.Name,
			&log.Calories,
			&log.Protein,
			&log.Carbs,
			&log.Fats,
			&log.MealType,
			&log.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}
