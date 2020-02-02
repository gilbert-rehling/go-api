package http

import (
	"fmt"
	"database/sql"
	"github.com/gilbert-rehling/go-api/models"
)

type PetHandler struct {
   Pets    models.PetStore
}

func getPets(ph *PetHandler, db *sql.DB) {
    var pets = {}

    pets, err := ph.findAllPets(db)

    // convert the response to string to ensure its not NULL
    response := string(pets[:])

    fmt.Println(response)

    fmt.Println("the end")
}