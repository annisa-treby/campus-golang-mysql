package repositories

import (
	"campus/models"
)

type FacultyRepository interface {
	GetAllFaculties() (*[]models.Faculty,error)
	InsertNewFaculty(faculty *models.Faculty)(*models.Faculty,error)
	UpdateFaculty(id string, faculty *models.Faculty)(*models.Faculty,error)
	GetFacultyById(id string)(*models.Faculty,error)
	DeleteFaculty(id string)error
	GetFacultyByName(name string)(*models.Faculty,error)
}
