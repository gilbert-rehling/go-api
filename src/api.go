package main

import (
    "time"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gilbert-rehling/go-api/db"
    "github.com/gilbert-rehling/go-api/http"
    "github.com/julienschmidt/httprouter"
)

var router = *httprouter.Router

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

    // Used during initial build: http.GetPets()
    // Default method uses native 'http'
    // Single Pet - call the handler
    // http.HandleFunc("/pet/{id}", http.GetPet())

    // Using imported router library
    // GET a single pet
    router.GET("/pet/:id", http.GetPet)

    // GET pet bu status
    router.GET("/pet/findByStatus", http.GetPetsByStatus)

    // init finished
    fmt.Println("init finished")
}

func main() {
    // acknowledge that main has been called
    fmt.Println("main triggered")

    // configuration - fixed IP or leave blank
    var server = "192.168.1.140"

    // configuration - service port !! required !!
    var port   = "8080"

    // set to listen on external facing IP address on DEV box
    log.Fatal(http.ListenAndServe(server + ":" + port, nil))
}
