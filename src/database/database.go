package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
    // DBCon is the DataBase connection handle
    DBCon *sql.DB
)