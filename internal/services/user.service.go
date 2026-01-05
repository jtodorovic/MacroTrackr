package services

import (
	"github.com/jtodorovic/macrotrackr/internal/db"
	"github.com/jtodorovic/macrotrackr/internal/models"
	"github.com/jtodorovic/macrotrackr/internal/security"
)

func NewUserResponse(u models.User) models.UserResponse {
	return models.UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		DOB:    u.DOB,
		Weight: u.Weight,
		Height: u.Height,
		Gender: u.Gender,
	}
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

func CreateUser(req models.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := security.HashPassword(req.Password)
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

	return &models.User{
		ID:     id,
		Name:   req.Name,
		Email:  req.Email,
		DOB:    req.DOB,
		Weight: req.Weight,
		Height: req.Height,
		Gender: req.Gender,
	}, nil
}
