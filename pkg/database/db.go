package database

import (
	"github.com/EmreAyberk/go-jwt-auth/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(database:3306)/go-jwt-auth?parseTime=true"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return DB, err
}
