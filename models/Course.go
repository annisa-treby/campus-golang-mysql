package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Course struct {
	BaseEntity
	Id string `json:"id"`
	Code string `json:"code"`
	Title string `json:"title"`
	MajorID string `json:"major_id"`
	Lecturer []*Lecturer `gorm:"many2many:lecturer_course;association_autoupdate:false;association_autoucreate:false"`
	Students []*Student `gorm:"many2many:student_course;association_autoupdate:false;association_autoucreate:false"`
}

func (c Course) TableName() string {
	return "course"
}
func (c *Course) BeforeCreate(db *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return err
}