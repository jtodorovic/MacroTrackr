package models

import (
	"time"

	"github.com/jtodorovic/macrotrackr/db"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	DOB       time.Time `json:"dob" binding:"required"`
	Weight    int32     `json:"weight" binding:"required"`
	Gender    string    `json:"gender" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var users = []User{}

func (u User) Save() error {
	query := `
INSERT INTO users(name, email, dob, weight, gender)
VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Name, u.Email, u.DOB, u.Weight, u.Gender)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.DOB, &user.Weight, &user.Gender)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
