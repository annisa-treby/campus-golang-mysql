package main

import (
	"campus/config"
	"campus/controllers"
	"campus/repositories"
	"campus/usecases"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	db := config.Connect()
	defer db.Close()
	db.LogMode(true)

	config.MigrateTable(db)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	facultyRepo := repositories.CreateFacultyRepository(db)
	facultyUseCase := usecases.CreateFacultyService(facultyRepo)
	controllers.CreateFacultyHandler(router, facultyUseCase)

	port := viper.GetString("port")

	err = router.Run(":"+port)
	if err != nil {
		log.Fatal(err)
	}
}
