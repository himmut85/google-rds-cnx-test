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

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "10.132.0.43"
	port     = 1433
	user     = "testuser"
	password = "test123"
)

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, user, password, port)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")
	defer conn.Close()
	stmt, err := conn.Prepare("select @@version")
	row := stmt.QueryRow()
	var result string

	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
