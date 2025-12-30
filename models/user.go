package models

import (
	"errors"
	"time"

	"github.com/jtodorovic/macrotrackr/db"
	"github.com/jtodorovic/macrotrackr/utils"
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

// DTOs
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

func NewUserResponse(u User) UserResponse {
	return UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		DOB:    u.DOB,
		Weight: u.Weight,
		Height: u.Height,
		Gender: u.Gender,
	}
}

type UserCredentials struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// func (u *User) ApplyUpdate(req UpdateUserRequest) error {
// 	if req.Name != nil {
// 		u.Name = *req.Name
// 	}
// 	if req.DOB != nil {
// 		u.DOB = *req.DOB
// 	}
// 	if req.Weight != nil {
// 		if *req.Weight <= 0 {
// 			return errors.New("weight must be positive")
// 		}
// 		u.Weight = *req.Weight
// 	}
// 	if req.Height != nil {
// 		if *req.Height <= 0 {
// 			return errors.New("height must be positive")
// 		}
// 		u.Height = *req.Height
// 	}
// 	if req.Gender != nil {
// 		u.Gender = *req.Gender
// 	}
// 	return nil
// }

func CreateUser(req CreateUserRequest) (*User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO users (name, email, password, dob, weight, height, gender)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(req.Name, req.Email, hashedPassword, req.DOB, req.Weight, req.Height, req.Gender)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:     id,
		Name:   req.Name,
		Email:  req.Email,
		DOB:    req.DOB,
		Weight: req.Weight,
		Height: req.Height,
		Gender: req.Gender,
	}, nil
}

func (uc *UserCredentials) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, uc.Email)

	var retrievedPassword string

	err := row.Scan(&uc.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	passwordIsValid := utils.ComparePassword(uc.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}

	return nil
}
