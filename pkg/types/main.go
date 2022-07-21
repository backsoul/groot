package types

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PayloadRegisterUser struct {
	FirstName string `validate:"required,min=3,max=32"`
	LastName  string `validate:"required,min=3,max=32"`
	Email     string `validate:"required,email,min=6,max=60"`
	Password  string `validate:"required,min=8,max=32"`
}

type PayloadLoginUser struct {
	Email    string `validate:"required,email,min=6,max=60"`
	Password string `validate:"required,min=8,max=32"`
}

type UserClaims struct {
	FirstName string
	LastName  string
	Email     string
	jwt.StandardClaims
}
