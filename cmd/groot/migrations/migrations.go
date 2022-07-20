package migrations

import (
	"github.com/backsoul/groot/pkg/types"
	"gorm.io/gorm"
)

func All(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&types.User{})
}
