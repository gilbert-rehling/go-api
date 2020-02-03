package http

import (
	"fmt"
	"github.com/gilbert-rehling/go-api/models"
)

func GetPets() {

    pets := models.FindAllPets()

    fmt.Println(pets)

    fmt.Println("http ended!")

}