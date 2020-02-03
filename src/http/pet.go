package http

// imports saved for later "log"
import (
	"fmt"
	"encoding/json"
	"github.com/gilbert-rehling/go-api/models"
)

// generic response function
func returnResponse(response []byte) {
    // just print for now
    fmt.Println(string(response))

    fmt.Println("response sent!")
}

//  GetPets calls FindAllPets
func GetPets() {
    results := models.FindAllPets()

    // prepare the response
    response, err := json.Marshal(results)
    if (err != nil) {
        // send empty response
        response, err = json.Marshal(nil)
        returnResponse(response)
    }

    // send response with data
    returnResponse(response)

    fmt.Println("GetPets ended!")

}

// GetPet call FindPetById
func GetPet() {
    var id int

    id = 1
    result := models.FindPetById( id )

    // prepare the response
    response, err := json.Marshal(result)
    if (err != nil) {
        // send empty response
        response, err = json.Marshal(nil)
        returnResponse(response)
    }

    // send response with data
    returnResponse(response)

    fmt.Println("GetPet ended!")

}