package database

import (
	"log"

	"github.com/szymon676/finance-tracker/server/src/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func SetupDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=finance port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.Income{}, &models.Spending{})
}
