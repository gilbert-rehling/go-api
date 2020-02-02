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
	"github.com/gilbert-rehling/go-api/models"
	"github.com/gilbert-rehling/go-api/http"
)

func init() {
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

    err = db.Ping()

    if err != nil {
        fmt.Println("ping error:", err)
        log.Panic("DB ping error: %s", err)
    }

    http.findAllPets(db)

}

func main() {
    // acknowledge the process has started
    fmt.Println("start")

}
