package models

import "time"

type (
	Laptop struct {
		Id              uint      `gorm:"primary_key" json:"id"`
		NamaLaptop      string    `json:"namaLaptop"`
		OperasiSistemID uint      `json:"osID"`
		BrandID         uint      `json:"brandID"`
		Layar           string    `json:"layar"`
		Hardisk         string    `json:"hardisk"`
		Ram             string    `json:"ram"`
		Grafis          string    `json:"grafis"`
		RatingID        uint      `json:"ratingID"`
		Harga           string    `json:"harga"`
		Image           string    `json:"image"`
		Prosesor        string    `json:"prosesor"`
		Detail          string    `json:"detail"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`

		Brand         Brand         `json:"-"`
		OperasiSistem OperasiSistem `json:"-"`
		Rating        Rating        `json:"-"`
	}
)
