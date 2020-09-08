package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Student struct {
	BaseEntity
	Id string `json:"id"`
	Name string `json:"student_name"`
	BirthDate time.Time `json:"birth_date"`
	Gender int `json:"gender"`
	Address string `json:"address"`
	Courses []*Course `gorm:"many2many:student_course;association_autoupdate:false;association_autoucreate:false"`
}

func (s Student) TableName() string {
	return "student"
}

func (s *Student) BeforeCreate(db *gorm.DB) (err error) {
	s.Id = uuid.New().String()
	return err
}
