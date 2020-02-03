package http

// imports saved for later "log"
import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gilbert-rehling/go-api/models"
)

// generic response function
func returnResponse(w http.ResponseWriter, response []byte) {
    // just print for now
    fmt.Println(string(response))

    fmt.Println("response sent!")
}

// GetPet call FindPetById
func GetPet(r *http.Request, ps httprouter.Params) {
    var id int

    id = ps.ByName("id")
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

//  GetPetsByStatus calls FindPetsByStatus
func GetPetsByStatus() {
    // retrieve the query parameter and call our model
    status := req.URL.Query().Get("status")
    results := models.FindPetsByStatus(status)

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
