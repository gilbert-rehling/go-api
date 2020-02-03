package http

import (
	"fmt"
	"database/sql"
	"github.com/gilbert-rehling/go-api/models"
)

func GetPets(db *sql.DB) {

    pet := models.FindAllPets("cat")

    fmt.Println(pet)

    fmt.Println("http ended!")

}