package models

import (
	"time"

	"github.com/jtodorovic/macrotrackr/db"
)

type FoodLog struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Name      string    `json:"name" binding:"required"`
	Calories  int32     `json:"calories" binding:"required"`
	Protein   int32     `json:"protein" binding:"required"`
	Carbs     int32     `json:"carbs" binding:"required"`
	Fats      int32     `json:"fats" binding:"required"`
	MealType  string    `json:"mealType" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (f *FoodLog) Save() error {
	query := `
	INSERT INTO food_logs(user_id, name, calories, protein, carbs, fats, meal_type)
	VALUES (?,?,?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(f.UserID, f.Name, f.Calories, f.Protein, f.Carbs, f.Fats, f.MealType)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	f.ID = id

	return err
}

func FindFoodLogByID(ID int64) (*FoodLog, error) {
	query := `SELECT * FROM food_logs WHERE id = ?`
	row := db.DB.QueryRow(query, ID)

	var foodLog FoodLog
	err := row.Scan(&foodLog.ID, &foodLog.UserID, &foodLog.Name, &foodLog.Calories, &foodLog.Protein, &foodLog.Carbs, &foodLog.Fats, &foodLog.MealType)
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

func FindFoodLogsByUserAndDate(userID int64, date time.Time) ([]FoodLog, error) {
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

	var logs []FoodLog

	for rows.Next() {
		var log FoodLog
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

func (f FoodLog) Update() error {
	query := `
	 UPDATE food_items
	 SET name = ?, calories = ?, protein = ?, carbs = ?, fats = ?, meal_type = ?
	 WHERE id = ?
	 `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(f.Name, f.Calories, f.Protein, f.Carbs, f.Fats, f.MealType)
	return err
}

func (f FoodLog) Delete() error {
	query := "DELETE FROM food_logs WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(f.ID)

	return err
}
