package main

// this router has issues !!  "github.com/julienschmidt/httprouter"
import (
    "time"
	"log"
	"net/http"
	"database/sql"
	"github.com/gilbert-rehling/go-api/db"
    "github.com/gilbert-rehling/go-api/handlers"
)

// var router *httprouter.Router

// initialises the application
func init() {
    // init started
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
        log.Panic("DB ping error: %s", err)
    }

    // Used during initial build: http.GetPets()
    // Default method uses native 'http'
    // Single Pet - call the handler
    // http.HandleFunc("/pet/{id}", http.GetPet())

    // Using imported router library
    // !! this flashy router wont work correctly for me at the moment
    //router := httprouter.New()

    // GET a single pet
    //router.GET("/pet/:var", handlers.GetPet)
    http.HandleFunc("/pet/find/id", handlers.GetPet)

    // GET pet by status
    //router.GET("/pet/:var/find", handlers.GetPetsByStatus)
    http.HandleFunc("/pet/find/by/status", handlers.GetPetsByStatus)

    // init finished
}

func main() {
    // main will load the listener
    // configuration - fixed IP or leave blank
    var server = "192.168.1.140"

    // configuration - service port !! required !!
    var port   = "8080"

    // Using imported router library
    // router := httprouter.New()

    // set to listen on external facing IP address on DEV box
    log.Fatal(http.ListenAndServe(server + ":" + port, nil))
}
