package models

import "time"

type (
	Brand struct {
		Id        uint      `gorm:"primary_key" json:"id"`
		NamaBrand string    `json:"namaBrand"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`

		Laptop []Laptop `json:"-"`
	}
)
