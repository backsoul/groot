package models

import (
	"github.com/backsoul/groot/internal/database"
	"github.com/backsoul/groot/pkg/types"
	"gorm.io/gorm"
)

func CreateUser(Uuid string, Name string, Email string, Provider string, Picture string) (tx *gorm.DB, err error) {
	tx = database.DB().Create(&types.User{
		ID:       Uuid,
		Name:     Name,
		Email:    Email,
		Provider: Provider,
		Picture:  Picture,
	})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}
