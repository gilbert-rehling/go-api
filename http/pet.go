package http

import (
	"fmt"
	"github.com/gilbert-rehling/go-api/models"
)

func GetPets() (string) {

    pet, err := models.FindAllPets("cat")

    fmt.Println("http ended!")

    return pet
}