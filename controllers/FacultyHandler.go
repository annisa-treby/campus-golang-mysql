package controllers

import (
	"campus/models"
	"campus/usecases"
	"campus/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FacultyHandler struct {
	facultyUseCase usecases.FacultyUseCase
}

func CreateFacultyHandler(r *gin.Engine, facultyUseCase usecases.FacultyUseCase)  {
	facultyHandler := FacultyHandler{facultyUseCase}

	r.GET("/faculties",facultyHandler.getAllFaculty)
	r.POST("/faculty", facultyHandler.insertNewFaculty)
}

func (f FacultyHandler) getAllFaculty(c *gin.Context) {
	faculties, err := f.facultyUseCase.GetAllFaculties()
	if err != nil {
		utils.HandleError(c, http.StatusBadGateway, err.Error())
	}
	utils.HandleSuccess(c, faculties)
}

func (f FacultyHandler) getFacultyByID(c *gin.Context) {
}

func (f FacultyHandler) insertNewFaculty(c *gin.Context) {
	var body models.Faculty

	err := c.Bind(&body)
	if err!=nil {
		utils.HandleError(c, http.StatusBadRequest,"Bad request body")
	}
	if body.Code == "" || body.FacultyName == "" {
		utils.HandleError(c, http.StatusBadRequest, "There's empty field")
	}

	faculty,err := f.facultyUseCase.InsertNewFaculty(&body)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError,"server error")
	}
	utils.HandleSuccess(c, faculty)
}