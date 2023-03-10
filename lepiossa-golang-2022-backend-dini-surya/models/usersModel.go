package models

import (
	"html"
	"lepiosa/utils/token"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Users struct {
		Id        uint      `gorm:"primary_key" json:"id"`
		Username  string    `gorm:"not null;unique" json:"userName"`
		Email     string    `gorm:"not null;unique" json:"email"`
		Password  string    `gorm:"not null;" json:"password"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Users{}

	err = db.Model(Users{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.Id)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Users) SaveUser(db *gorm.DB) (*Users, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Users{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}
