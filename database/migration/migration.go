package migration

import (
	"filmes-crud/model"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(model.Movie{})

}
