package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Faculty struct {
	BaseEntity
	ID string `gorm:"primary_key"`
	Code string `json:"code"`
	FacultyName string `json:"faculty_name"`
	Majors []Major `json:"majors"`
}

func (f Faculty) TableName() string {
	return "faculty"
}

func (f *Faculty) BeforeCreate(db *gorm.DB) (err error) {
	f.ID = uuid.New().String()
	return err
}
