package models

import "time"

type Summary struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	UserID          string    `json:"userId"`
	Date            time.Time `json:"date"`
	CurrentCalories int32     `json:"currentCalories"`
	CurrentProtein  int32     `json:"currentProtein"`
	CurrentCarbs    int32     `json:"currentCarbs"`
	CurrentFats     int32     `json:"currentFats"`
	// should i put goal here also ?
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}
