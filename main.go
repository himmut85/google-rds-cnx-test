package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHostName := "34.140.158.202"
	dbUserName := "postgres"
	dbUserPass := "Test123"
	dbName := "postgres"
	dbPort := "5432"
	dbSslMode := "disable"
	dbTimeZone := "UTC"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHostName, dbUserName, dbUserPass, dbName, dbPort, dbSslMode, dbTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	fmt.Println(db)
	
	if err != nil {
		panic("DB connection could not be established!")
	}

	fmt.Println(db)

}
