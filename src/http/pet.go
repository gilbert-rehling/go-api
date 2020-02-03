package http

import (
	"fmt"
	"database/sql"
	"github.com/gilbert-rehling/go-api/models"
)

func GetPets(db *sql.DB) {

    var pets string

    pets = models.FindAllPets(db)

    fmt.Println(pets)

    fmt.Println("http ended!")

}