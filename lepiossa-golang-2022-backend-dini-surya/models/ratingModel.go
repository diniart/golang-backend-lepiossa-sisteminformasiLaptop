package models

import "time"

type (
	Rating struct {
		Id        uint      `gorm:"primary_key" json:"id"`
		Rating    int       `json:"rating"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`

		Laptop []Laptop `json:"-"`
	}
)
