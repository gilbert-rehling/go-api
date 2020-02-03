package models

// spare imports ->  "database/sql" "time"
import (
    "fmt"
    "github.com/gilbert-rehling/go-api/db"
)

// Pet defines the 'pet' properties structure
// For the time being I will use string for the dates instead of: time.Time
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

// FindPetById returns a single pet by id column
func FindPetById( id int) (Pet) {
    // output container
    pet := Pet{}

    // generate query
    var stmt = "SELECT * FROM `pet` WHERE `id` = ? ORDER BY `id` DESC"

    // run the query
    err := db.Conn.QueryRow(stmt, id).Scan(&pet.ID, &pet.Category, &pet.Name, &pet.PhotoUrls, &pet.Tags, &pet.Status, &pet.Updated, &pet.Created)
    if err != nil {
         fmt.Printf(err.Error()) // implement proper error handling instead of panic
    }

    return pet
}

// FindPetsByStatus returns all pets matching the status parameter
func FindPetsByStatus(status string) ([]Pet) {
    // the output container
    var pets []Pet

    // generate query
    var stmt = "SELECT `id`, `status` FROM `pet` WHERE `status` = ? ORDER BY `status` ASC"

    // run the query
    rows, err := db.Conn.Query(stmt, status)
    if err != nil {
         panic(err.Error()) // implement proper error handling instead of panic
    }

    // iterate the resulting rows
    // ToDO: this was the simplest and shortest (least amount of code) method to fetch all the results into a structure
    for rows.Next() {

        pet := Pet{}
        // get RawBytes from data
        err = rows.Scan(&pet.ID, &pet.Status)
        if err != nil {
            panic(err.Error()) // implement proper error handling instead of panic
        }
        // build up the output container
        pets = append(pets, pet)

    }
    if err = rows.Err(); err != nil {
        panic(err.Error()) // implement proper error handling instead of panic
    }

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
         panic(err.Error()) // implement proper error handling instead of panic
    }

    // iterate the resulting rows
    // ToDO: this was the simplest and shortest (least amount of code) method to fetch all the results into a structure
    for rows.Next() {

        pet := Pet{}
        // get RawBytes from data
        err = rows.Scan(&pet.ID, &pet.Category, &pet.Name, &pet.PhotoUrls, &pet.Tags, &pet.Status, &pet.Updated, &pet.Created)
        if err != nil {
            panic(err.Error()) // implement proper error handling instead of panic
        }
        // build up the output container
        pets = append(pets, pet)

    }
    if err = rows.Err(); err != nil {
        panic(err.Error()) // implement proper error handling instead of panic
    }

    return pets
}