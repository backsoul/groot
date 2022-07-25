package types

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Provider  string
	Name      string
	Picture   string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserClaims struct {
	Picture string
	Email   string
	Name    string
	jwt.StandardClaims
}
