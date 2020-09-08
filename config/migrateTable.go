package config

import (
	"campus/models"
	"github.com/jinzhu/gorm"
)

func MigrateTable(db *gorm.DB)  {
	db.AutoMigrate(&models.Faculty{},
		&models.Lecturer{},
		&models.Course{},
		&models.Score{},
		&models.Major{},
		&models.Student{})
}
