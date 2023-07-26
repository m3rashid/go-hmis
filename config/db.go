package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() {
	dsn := "host=localhost user=genos password=genos dbname=hmis port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
