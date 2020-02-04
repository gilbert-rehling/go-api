package main

// this router has issues !!  "github.com/julienschmidt/httprouter"
import (
    "time"
	"log"
	"net/http"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"github.com/gilbert-rehling/go-api/db"
    "github.com/gilbert-rehling/go-api/handlers"
)

 //var router *httprouter.Router

// initialises the application
//func init() {
    // init started

    // init finished
//}

func main() {
    // main will load the listener
    // configuration - fixed IP or leave blank
    var server = "192.168.1.140"

    // configuration - service port !! required !!
    var port   = "8080"

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

    // Using imported router library
    // !! this flashy router wont work correctly for me at the moment
    router := httprouter.New()

    // This router isn't working for me !! keep getting 404's
    // But this 'wildcard' route works !!
    router.GET("/pet/*segment", handlers.GetSwitch)

    // Send to the POST switch
    router.POST("/pet/*segment", handlers.PostSwitch)

    // There is only 1 DELETE route
    router.DELETE("/pet/*segment", handlers.DeletePet)

    // There is only 1 PUT route
    // This needs to share the UpdatePet handler
    router.PUT("/pet/*segment", handlers.PutSwitch)

    // set to listen on external facing IP address on DEV box
    log.Fatal(http.ListenAndServe(server + ":" + port, router))
}
