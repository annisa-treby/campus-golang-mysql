package utils

import (
	"campus/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleSuccess(c *gin.Context, data interface{}) {
	responseData := models.Response{
		Status: "200",
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, responseData)
}

func HandleError(c *gin.Context, status int, message string) {
	responseData := models.Response{
		Status: strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responseData)
}
