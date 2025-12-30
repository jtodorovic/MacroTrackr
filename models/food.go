package models

import "time"

type Food struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"userId" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Calories  int32     `json:"calories" binding:"required"`
	Protein   int32     `json:"protein" binding:"required"`
	Carbs     int32     `json:"carbs" binding:"required"`
	Fats      int32     `json:"fats" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}
