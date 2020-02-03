package main

import (
    "time"
	"fmt"
	"log"
	"github.com/gilbert-rehling/go-api/database"
    "github.com/gilbert-rehling/go-api/http"
)

func init() {

    fmt.Println("init start")

    var er error

    // create the DB link !! todo: using hard coded credentials for testing !!
    database.DBCon, err = sql.Open("mysql", "petapi:petapi2020@tcp(localhost:3306)/petapi")

    if err != nil {
        log.Panic("DB connection error: %s", err)
    }

    // To avoid error due to connection being closed on the server
    database.DBCon.SetConnMaxLifetime(14000 * time.Second)

    // check that the db link works
    err = database.DBCon.Ping()

    if err != nil {
        fmt.Println("ping error:", err)
        log.Panic("DB ping error: %s", err)
    }

    // call the handler - pass the DB connection for injection
    http.GetPets()

    fmt.Println("init finished")

}

func main() {
    // acknowledge the process has started
    fmt.Println("main triggered")
}
