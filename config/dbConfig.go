package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

type DataSourceName struct {
	DbUser string
	DbPass string
	DbHost string
	DbName string
	DbPort string
	DbEngine string
}

func Connect() *gorm.DB {
	//viper.AutomaticEnv()
	//viper.SetConfigFile(".env")
	//
	//err := viper.ReadInConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}

	DB := DataSourceName{
		DbUser: viper.GetString("DB_USER"),
		DbPass: viper.GetString("DB_PASS"),
		DbHost: viper.GetString("DB_HOST"),
		DbName: viper.GetString("DB_NAME"),
		DbPort: viper.GetString("DB_PORT"),
		DbEngine: viper.GetString("DB_ENGINE"),
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DB.DbUser,DB.DbPass,DB.DbHost,DB.DbPort,DB.DbName)

	db,err := gorm.Open(DB.DbEngine,conn)
	if err!=nil {
		log.Fatal(err)
	}
	return db
}
