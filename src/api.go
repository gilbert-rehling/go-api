package main

import (
    "time"
	"fmt"
	"log"
	"database/sql"
	"github.com/gilbert-rehling/go-api/db"
    "github.com/gilbert-rehling/go-api/http"
)

// initialises the application
func init() {
    // init started
    fmt.Println("init start")

    var err error

    // create the DB link !! todo: using hard coded credentials for testing !!
    db.Conn, err = sql.Open("mysql", "petapi:petapi2020@tcp(localhost:3306)/petapi")
    if err != nil {
        log.Panic("DB connection error: %s", err)
    }

    // To avoid error due to connection being closed on the server
    db.Conn.SetConnMaxLifetime(14000 * time.Second)

    // check that the db link works
    err = db.Conn.Ping()
    if err != nil {
        fmt.Println("ping error:", err)
        log.Panic("DB ping error: %s", err)
    }

    // All Pets - call the handler
    //http.GetPets()

    // Single Pet - call the handler
    http.GetPet()

    // init finished
    fmt.Println("init finished")
}

func main() {
    // acknowledge that main has been called
    fmt.Println("main triggered")
}
