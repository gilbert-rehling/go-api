package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
    // Conn is the DataBase connection handle
    Conn *sql.DB
)