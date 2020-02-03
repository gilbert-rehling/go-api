package main

import (
    "time"
	"fmt"
	"log"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gilbert-rehling/go-api/http"
)

func init() {

    fmt.Println("init start")

    // create the DB link !! todo: using hard coded credentials for testing !!
    db, err := sql.Open("mysql", "petapi:petapi2020@tcp(localhost:3306)/petapi")

    if err != nil {
        log.Panic("DB connection error: %s", err)
    }

    // To avoid error due to connection being closed on the server
    db.SetConnMaxLifetime(14000 * time.Second)

    // check that the db link works
    err = db.Ping()

    if err != nil {
        fmt.Println("ping error:", err)
        log.Panic("DB ping error: %s", err)
    }

    // call the handler - pass the DB connection for injection
    http.GetPets(db)

    fmt.Println("init finished")

}

func main() {
    // acknowledge the process has started
    fmt.Println("main triggered")
}
