package http

import (
	"fmt"
	"github.com/gilbert-rehling/go-api/models"
)

func GetPets() (string) {

    pet := models.FindAllPets("cat")

    fmt.Println("http ended!")

    return pet
}