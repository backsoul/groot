package models

import (
	"github.com/backsoul/groot/internal/database"
	"github.com/backsoul/groot/pkg/types"
	"gorm.io/gorm"
)

func CreateUser(FirstName string, LastName string, Email string, Password string) (tx *gorm.DB, err error) {
	tx = database.DB().Create(&types.User{
		FirstName: FirstName,
		LastName:  FirstName,
		Email:     Email,
		Password:  Password,
	})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}
