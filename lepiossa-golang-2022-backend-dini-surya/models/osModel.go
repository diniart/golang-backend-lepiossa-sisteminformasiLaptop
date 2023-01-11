package models

import "time"

type (
	OperasiSistem struct {
		Id        uint      `gorm:"primary_key" json:"id"`
		NamaOs    string    `json:"namaOs"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`

		Laptop []Laptop `json:"-"`
	}
)
