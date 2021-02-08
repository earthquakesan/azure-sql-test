// Go connection Sample Code:
package main
import _ "github.com/denisenkom/go-mssqldb"
import (
	"database/sql"
	"context"
	"log"
	"fmt"
	"os"
	"strconv"
)

var db *sql.DB

func main() {
	var err error
		
	var user = os.Getenv("SQL_USERNAME")
	var password = os.Getenv("SQL_PASSWORD")
	var server = os.Getenv("SQL_HOSTNAME")
	var database = os.Getenv("SQL_DATABASE")
	var port = 1433
	port, err = strconv.Atoi(os.Getenv("SQL_PORT"))

	if err != nil {
		log.Fatal(err.Error())
	}

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
	os.Exit(0)
}