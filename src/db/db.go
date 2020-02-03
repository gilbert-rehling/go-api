package db

// importing the /mysql driver here seems to work nicely
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// import this package wherever you need DB connectivity...
// Usage:  db.Conn.Open() -> db.conn.Query()
var (
    // Conn is the DataBase connection handle
    Conn *sql.DB
)