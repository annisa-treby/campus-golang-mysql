package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Score struct {
	Id       string `json:"id"`
	Alphabet string `json:"alphabet"`
}

func (s Score) TableName() string {
	return "score"
}

func (s *Score) BeforeCreate(db *gorm.DB) (err error) {
	s.Id = uuid.New().String()
	return err
}