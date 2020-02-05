package models

// In order to maintain best practice:
// 1) I am minimising the imports to one local DB package
// 2) We are always return the expected structure back to the handlers
// !! These models are specific to the DB table and route specifications provided for the project scope !!
// spare imports ->  "database/sql" "time"
import (
    "database/sql"
    "github.com/gilbert-rehling/go-api/db"
)

// Define out structures to return to the handlers
// For the time being I will use string for the dates instead of: time.Time
// Pet defines the 'pet' properties structure
type Pet struct {
	ID        int       `json:"id"`
	Category  int       `json:"category_id"`
	Name      string    `json:"name"`
	PhotoUrls string    `json:"photo_urls"`
	Tags      string    `json:"tags"`
	Status    string    `json:"status"`
	Updated   string    `json:"updated_at"`
	Created   string    `json:"created_at"`
}

// PetByStatus defines the 'pet' properties structure
type PetByStatus struct {
    ID        int       `json:"id"`
    Status    string    `json:"status"`
}

// FindPetById returns a single pet by id column
func FindPetById( id string) map[string]interface{} {
    // output container
    var pet = make(map[string]interface{})

    // generate query
    var stmt = "SELECT * FROM `pet` WHERE `id` = ? ORDER BY `id` DESC"

    // run the query
    // I have now changed this from QueryRow() to Query() so that I can use dynamic mapping
    // Word on the street is that structures are faster -
    // but it means I'd have to create structures for each different model response (Urghh!!)
    rows, _ := db.Conn.Query(stmt, id)

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        // just return the empty structure for now
        // add some logging etc later
        return pet
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument,
    // so we must copy the references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    // iterate the rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            // just return the empty structure for now
            // add some logging etc later
            return pet
        }
        var value string
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            pet[columns[i]] = value
        }
    }
    if err = rows.Err(); err != nil {
        // just return the empty structure for now
        // add some logging etc later
        return pet
    }

    // return to handler
    return pet
}

// FindPetsByStatus returns all pets matching the status parameter
func FindPetsByStatus(status string) map[int]interface{} {
    // the output container
    var pets = make(map[int]interface{})

    // generate query
    var stmt = "SELECT `id`, `status` FROM `pet` WHERE `status` = ? ORDER BY `status` ASC"

    // run the query
    rows, err := db.Conn.Query(stmt, status)
    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        // just return the empty structure for now
        // add some logging etc later
        return pets
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument,
    // so we must copy the references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    // iterate the rows
    var x = 0;
    for rows.Next() {
        pet := make(map[string]interface{})
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            // just return the empty structure for now
            // add some logging etc later
            return pets
        }
        var value string
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            pet[columns[i]] = value
        }
        pets[x] = pet
        x++
    }
    if err = rows.Err(); err != nil {
        // just return the empty structure for now
        // add some logging etc later
        return pets
    }

    // return to handler
    return pets
}

// FindAllPets returns all pets from the pet table
func FindAllPets() ([]Pet) {
    // the output container
    var pets []Pet

    // generate query
    var stmt = "SELECT * FROM `pet` ORDER BY `id` DESC"

    // run the query
    rows, err := db.Conn.Query(stmt)
    if err != nil {
         // just return the empty structure for now
         // add some logging etc later
         return pets
    }

    // iterate the resulting rows
    // ToDO: this was the simplest and shortest (least amount of code) method to fetch all the results into a structure
    for rows.Next() {

        pet := Pet{}
        // get RawBytes from data
        err = rows.Scan(&pet.ID, &pet.Category, &pet.Name, &pet.PhotoUrls, &pet.Tags, &pet.Status, &pet.Updated, &pet.Created)
        if err != nil {
            // just return the empty structure for now
            // add some logging etc later
            return pets
        }
        // build up the output container
        pets = append(pets, pet)

    }
    if err = rows.Err(); err != nil {
        // just return the empty structure for now
        // add some logging etc later
        return pets
    }

    return pets
}