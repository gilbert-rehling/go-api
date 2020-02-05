package models

// In order to maintain best practice:
// 1) I am minimising the imports to one local DB package
// 2) We are always return the expected structure back to the handlers
// !! These models are specific to the DB table and route specifications provided for the project scope !!
// spare imports ->  "database/sql" "time"
import (
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
func FindPetById( id string) (Pet) {
    // output container
    pet := Pet{}

    // generate query
    var stmt = "SELECT * FROM `pet` WHERE `id` = ? ORDER BY `id` DESC"

    // run the query
    err := db.Conn.QueryRow(stmt, id).Scan(&pet.ID, &pet.Category, &pet.Name, &pet.PhotoUrls, &pet.Tags, &pet.Status, &pet.Updated, &pet.Created)
    if err != nil {
         // The error here would be 'no result' to run Scan with (sql.ErrNoRows)
         // Just return the empty pet
         return pet
    }

    // return to handler
    return pet
}

// FindPetsByStatus returns all pets matching the status parameter
func FindPetsByStatus(status string) ([]PetByStatus) {
    // the output container
    var pets []PetByStatus

    // generate query
    var stmt = "SELECT `id`, `status` FROM `pet` WHERE `status` = ? ORDER BY `status` ASC"

    // run the query
    rows, err := db.Conn.Query(stmt, status)
    if err != nil {
        // just return the empty structure for now
        // add some logging etc later
        return pets
    }

    // iterate the resulting rows
    // ToDO: this was the simplest and shortest (least amount of code) method to iterate the result set for restructuring
    for rows.Next() {

        pet := PetByStatus{}
        // get RawBytes from data
        err = rows.Scan(&pet.ID, &pet.Status)
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