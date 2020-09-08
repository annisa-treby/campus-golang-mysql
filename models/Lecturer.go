package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Lecturer struct {
	BaseEntity
	Id string `json:"id"`
	Name string `json:"name"`
	Gender int `json:"gender"`
	BirthDate string `json:"birth_date"`
	Address string `json:"address"`
	Courses []*Course `gorm:"many2many:lecturer_course;association_autoupdate:false;association_autoucreate:false"`
}

func (l Lecturer) TableName() string{
	return "lecturer"
}

func (l *Lecturer) BeforeCreate(db *gorm.DB) (err error) {
	l.Id = uuid.New().String()
	return err
}
