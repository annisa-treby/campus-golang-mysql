package usecases

import (
	"campus/models"
	"campus/repositories"
)

type FacultyServiceImpl struct {
	facultyRepo repositories.FacultyRepository
}

func CreateFacultyService(facultyRepo repositories.FacultyRepository) FacultyUseCase {
	return &FacultyServiceImpl{facultyRepo}
}

func (f FacultyServiceImpl) GetAllFaculties() (*[]models.Faculty, error) {
	return f.facultyRepo.GetAllFaculties()
}

func (f FacultyServiceImpl) InsertNewFaculty(faculty *models.Faculty) (*models.Faculty, error) {
	return f.facultyRepo.InsertNewFaculty(faculty)
}

func (f FacultyServiceImpl) UpdateFaculty(id string, faculty *models.Faculty) (*models.Faculty, error) {
	return f.facultyRepo.UpdateFaculty(id, faculty)
}

func (f FacultyServiceImpl) GetFacultyById(id string) (*models.Faculty, error) {
	return f.facultyRepo.GetFacultyById(id)
}

func (f FacultyServiceImpl) DeleteFaculty(id string) error {
	return f.facultyRepo.DeleteFaculty(id)
}

func (f FacultyServiceImpl) GetFacultyByName(name string) (*models.Faculty, error) {
	return f.facultyRepo.GetFacultyByName(name)
}
