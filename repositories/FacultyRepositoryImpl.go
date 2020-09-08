package repositories

import (
	"campus/models"
	"github.com/jinzhu/gorm"
	"log"
)

type FacultyRepoImpl struct {
	db *gorm.DB
}

func CreateFacultyRepository(db *gorm.DB) FacultyRepository {
	return &FacultyRepoImpl{db}
}

func (f FacultyRepoImpl) GetAllFaculties() (*[]models.Faculty, error) {
	var faculties []models.Faculty
	err := f.db.Set("gorm:auto_preload",true).Find(&faculties).Error
	if err!=nil {
		log.Fatal(err)
	}
	return &faculties,err
}

func (f FacultyRepoImpl) InsertNewFaculty(faculty *models.Faculty) (*models.Faculty, error) {
	err := f.db.Create(&faculty).Error
	if err!=nil {
		log.Fatal(err)
	}
	return faculty,err
}

func (f FacultyRepoImpl) UpdateFaculty(id string, faculty *models.Faculty) (*models.Faculty,error){
	err := f.db.Model(&faculty).Where("id=?",id).Update(faculty).Error
	if err!=nil {
		log.Fatal(err)
	}
	return faculty,err
}

func (f FacultyRepoImpl) GetFacultyById(id string) (*models.Faculty, error) {
	var faculty models.Faculty
	err := f.db.Set("gorm:auto_preload",true).First(&faculty, id).Error
	if err!=nil {
		log.Fatal(err)
	}
	return &faculty,err
}

func (f FacultyRepoImpl) DeleteFaculty(id string) error {
	faculties := models.Faculty{}
	err:=f.db.Table("faculty").Where("id=?",id).First(&faculties).Delete(&faculties).Error
	if err!=nil {
		log.Fatal(err)
	}
	return err
}

func (f FacultyRepoImpl) GetFacultyByName(name string) (*models.Faculty, error) {
	var faculty models.Faculty
	err := f.db.First(&faculty, "faculty_name LIKE ?", name).Error
	if err!=nil {
		log.Fatal(err)
	}
	return &faculty,err
}

