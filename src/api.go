package main

// imports
import (
    "bufio"
    "database/sql"
	"encoding/json"
	"net"
	"net/url"
	"net/http"
	"log"
	"regexp"
)

// container for the widgets
type Request struct {
	ID string
	Name string
	Headers map[string]string
	Params map[string]string
}

const port 3306

var db *sql.DB

func call_result(conn net.Conn, request *Request) {

}

func dbExec(stmt string, parameters ...interface{}) {

	_, err := db.Exec(stmt, parameters...)

	if err != nil {
		log.Print("Failed to execute SQL statement: %s; stmt: %s", err, stmt)
	}
}
