package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
        "net/http"
        "os"
	
)

func main() {
	dbHostName := "172.21.0.3"
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
	
	
	port := os.Getenv("PORT")
  	if port == "" {
        	port = "8080"
	}	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	

}

