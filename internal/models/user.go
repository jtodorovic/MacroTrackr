package models

import (
	"errors"
	"time"

	"github.com/jtodorovic/macrotrackr/internal/db"
	"github.com/jtodorovic/macrotrackr/internal/security"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	DOB       time.Time
	Weight    int32
	Height    int32
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserRequest struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=5"`
	DOB      time.Time `json:"dob" binding:"required"`
	Weight   int32     `json:"weight" binding:"required,gt=0"`
	Height   int32     `json:"height" binding:"required,gt=0"`
	Gender   string    `json:"gender" binding:"required,oneof=male female"`
}

// type UpdateUserRequest struct {
// 	Name   *string    `json:"name,omitempty"`
// 	DOB    *time.Time `json:"dob,omitempty"`
// 	Weight *int32     `json:"weight,omitempty"`
// 	Height *int32     `json:"height,omitempty"`
// 	Gender *string    `json:"gender,omitempty"`
// }

type UserResponse struct {
	ID     int64     `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	DOB    time.Time `json:"dob"`
	Weight int32     `json:"weight"`
	Height int32     `json:"height"`
	Gender string    `json:"gender"`
}

type UserCredentials struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (uc *UserCredentials) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, uc.Email)

	var retrievedPassword string

	err := row.Scan(&uc.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	passwordIsValid := security.ComparePassword(uc.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}

	return nil
}
