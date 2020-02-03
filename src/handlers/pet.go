package handlers

// imports saved for later "log" "github.com/julienschmidt/httprouter"
import (
	"net/http"
	"encoding/json"
	"github.com/gilbert-rehling/go-api/models"
)

type Data struct {
    Error   bool    `json:"error"`
    Message string  `json:"message"`
}

type Error struct {
    Code    int     `json:"code"`
    Type    string  `json:"type"`
    Message string  `json:"message:"`
    Data    Data    `json:"data"`
}

// GetPet call FindPetById - , ps httprouter.Params
func GetPet(w http.ResponseWriter, r *http.Request) {
    // the id should be passed within the URI as a segment
    //id := ps.ByName("var")
    id := r.URL.Query().Get("id")
    result := models.FindPetById( id )

    // Success response structure
    // Different Data structure for responses
    type Response struct {
        Code    int     `json:"code"`
        Type    string  `json:"type"`
        Message string  `json:"message"`
        Data    models.Pet  `json:"data"`
    }

    var rtn Response
    rtn.Code = 200
    rtn.Type = "Success"
    rtn.Message = "Operation executed successfully"
    rtn.Data = result

    // response as JSON
    response, err := json.Marshal(rtn)
    if (err != nil) {
        // send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

//  GetPetsByStatus calls FindPetsByStatus - , ps httprouter.Params
func GetPetsByStatus(w http.ResponseWriter, r *http.Request) {
    // retrieve the query parameter and call our model
    status := r.URL.Query().Get("status")
    //status := ps.ByName("var")
    if (status == "") {
        // send error response
        var d Data
        d.Error = true
        d.Message = "status missing"

        var er Error
        er.Code = 200
        er.Type = "Success"
        er.Message = "Operation executed successfully"
        er.Data = d

        response, err := json.Marshal(er)
        if (err != nil) {
            // send error response
            w.Header().Set("Content-Type", "application/json")
            w.Write(response)
        }
    }

    // run query on model
    results := models.FindPetsByStatus(status)

    // Success response structure
    // Different Data structure for responses
    type Response struct {
        Code    int     `json:"code"`
        Type    string  `json:"type"`
        Message string  `json:"message"`
        Data    []models.PetByStatus  `json:"data"`
    }

    var rtn Response
    rtn.Code = 200
    rtn.Type = "Success"
    rtn.Message = "Operation executed successfully"
    rtn.Data = results

    // response as JSON
    response, err := json.Marshal(rtn)
    if (err != nil) {
        // send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}
