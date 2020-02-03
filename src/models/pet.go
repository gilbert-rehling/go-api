package models

import (
    "fmt"
    "time"
    "database/sql"
    "github.com/gilbert-rehling/go-api/db"
)

// Pet defines the pet properties
type Pet struct {
	ID        int       `json:"id"`
	Category  int       `json:"category_id"`
	Name      string    `json:"name"`
	PhotoUrls string    `json:"photo_urls"`
	Tags      string    `json:"tags"`
	Status    string    `json:"status"`
	Updated   time.Time `json:"updated_at"`
	Created   time.Time `json:"created_at"`
}

func FindAllPets() ([]Pet) {
    //var pets []Pet

    // generate query
    var stmt = "SELECT * FROM `pet` ORDER BY `id` DESC"

    // run the query
    rows, err := db.Conn.Query(stmt)
    if err != nil {
        return ''
    }

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    pets := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(pets...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            fmt.Println(columns[i], ": ", value)
        }
        fmt.Println("-----------------------------------")
    }
    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    return pets
}