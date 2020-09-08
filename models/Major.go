package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Major struct {
	BaseEntity
	Id string `json:"id"`
	Code string `json:"code"`
	MajorName string `json:"major_name"`
	FacultyID string `json:"faculty_id"`
	Courses []Course `json:"courses"`
}

func (m Major) TableName() string {
	return "major"
}
func (m *Major) BeforeCreate(db *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return err
}