package main

// imports - "regexp" "net/http" "encoding/json" "net"
import (
    "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"log"
	"log/syslog"
	"os"
	"fmt"
)

/*
 * Tag... - a very simple struct
 */
type Tag struct {
    Id   string `json:"id"`
    Name string `json:"name"`
    Status string `json:"status"`
}

var db *sql.DB

const port = "3306"

// params - db *sql.DB
func getResults() {

    var tag Tag

    err := db.QueryRow("SELECT `id`, `name`, `status` FROM `pet` WHERE `id` = ?", 1).Scan(&tag.Id, &tag.Name, &tag.Status)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    fmt.Println("id value: " + tag.Id)

    fmt.Println("name value: " + tag.Name)

    fmt.Println("status value: " + tag.Status)

}

func main() {

    // acknowledge the process has started
    fmt.Println("start")

    // handle some local logging
    logwriter, err := syslog.New(syslog.LOG_WARNING, "go API running")

    if err == nil {
    	log.SetOutput(logwriter)
    }

    var dir string

    // get the current directory
    dir, err = os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(dir)

    db, dbErr := sql.Open("mysql", "petapi:petapi2020@tcp(localhost:3306)/petapi")

    if dbErr != nil {
    	log.Panic("DB connection error: %s", dbErr)
    }

    // To avoid error due to connection being closed in the server
    db.SetConnMaxLifetime(14000 * time.Second)

    // defer db.Close()

    err = db.Ping()

    if err != nil {
        fmt.Println("ping error:", err)
        log.Panic("DB ping error: %s", err)
    }

    //listener, err := net.Listen("tcp", ":"+port)

    //if err != nil {
        //log.Panic("Cannot bind to port")
    //}

    //conn, err := listener.Accept() // this blocks until connection or error

    //if err != nil {
        //log.Panic("Cannot accept connection")
    //}

    getResults()

}
