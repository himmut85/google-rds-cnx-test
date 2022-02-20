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
	_ "github.com/denisenkom/go-mssqldb"
)



func main() {
	
	var n_tables int

   println (sql.Drivers())

   // URL connection string formats
   //    sqlserver://sa:mypass@localhost?database=master&connection+timeout=30         // username=sa, password=mypass.
   //    sqlserver://sa:my%7Bpass@somehost?connection+timeout=30                       // password is "my{pass"
   // note: pwd is "myP@55w0rd"
   connectString := "sqlserver://SBM:myP%4055w0rd@VM17:1433?database=AE&connection+timeout=30"
   println("Connection string=" , connectString )

   println("open connection")
   db, err := sql.Open("mssql", connectString)
   defer db.Close()
   println ("Open Error:" , err)
   if err != nil {
      log.Fatal(err)
   }else {println ("SQL CONNECTED")}

   println("count records in TS_TABLES & scan")
   err = db.QueryRow("Select count(*) from ts_tables").Scan(&n_tables)
   if err != nil {
      log.Fatal(err)
   }
   println ("count of tables" , n_tables)

   println("closing connection")
   db.Close()
	
	
	port := os.Getenv("PORT")
  	if port == "" {
        	port = "8080"
	}	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
