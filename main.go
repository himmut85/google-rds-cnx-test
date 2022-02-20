/* package main

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
*/



package main

import (
	"database/sql"
	"fmt"
	"log"
        "net/http"
        "os"
	import _ "github.com/go-sql-driver/mysql"
)

var (
	server   = "10.132.0.4312"
	port     = "1432"
	user     = "testuser"
	password = "test123"
	
)

func main() {
	
	
	
	
	connString := fmt.Sprintf("server=%s;user=%s;password=%s;port=%s",
		server, user, password, port)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("SQL Open connection failed:", err.Error())
	}else {
		fmt.Printf("sql Connected!\n")
	}
	defer conn.Close()
	
	
	port := os.Getenv("PORT")
  	if port == "" {
        	port = "8080"
	}	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
