package handlers

// imports saved for later "log" "github.com/julienschmidt/httprouter"   "fmt"
import (
    "strings"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
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

// GetSegments determines our function and possible value segments from the URL
// We are assuming a well defined URI pattern, like ~.com/root/function/value,
// but in the switch statement I am handling the routes defined in the swagger pet store as an exercise only
func BuildSegments(r *http.Request) map[string]string {
    // get the url segments
    uriSegments := strings.Split(r.URL.Path, "/")

    segments := make(map[string]string)

    // if we are here the the uri root must be /pet
    segments["root"] = uriSegments[1]

    // predefine the map elements as empty
    segments["function"] = ""
    segments["value"] = ""

    // get the 'function'
    if (len(uriSegments) > 2) {
        segments["function"] = uriSegments[2]
    }

    // get the 'value' - could be int or string
    if (len(uriSegments) > 3) {
        segments["value"] = uriSegments[3]
    }

    return segments
}

// GetSwitch we use this to determine which handler to run based on the URL structure
// The httprouter package was causing errors, but the wildcard loader works fine
// We will create a switch function for each HTTP request type so that they dont get to crowded
// We are also defining methods that match up with the swagger pet store project
func GetSwitch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // get the segments
    var segments = BuildSegments(r)

    switch segments["function"] {
        case "findByStatus":
            GetPetsByStatus(w, r, segments["value"])

        default:
            if (segments["value"] == "") {
                segments["value"] = segments["function"]
            }
            GetPet(w, r, segments["value"])
    }
}

// PostSwitch handles dispatching the 3 POST routes to the correct handler
// ToDo: !! Currently the routing package is not set to handle ~.com/pet -> it needs to be ~.com/pet/ with a trailing slash
func PostSwitch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // ToDo: !! this block need to be updated !!
    // get the segments
    var segments = BuildSegments(r)

    switch segments["function"] {
        case "":
            PostPet(w, r, segments["value"])

        default:
            if (segments["value"] == "uploadImage") {
                UploadImage(w, r, segments["function"])
            } else {
                UpdatePet(w, r, ps, segments["function"])
            }
    }
}

// PutSwitch this enables us to share the update function between PUT & POST
func PutSwitch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    UpdatePet(w, r, ps, "")
}

// GetPet call FindPetById
func GetPet(w http.ResponseWriter, r *http.Request, id string) {
    // the id should be passed within the URI as a segment
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

//  GetPetsByStatus calls FindPetsByStatus
func GetPetsByStatus(w http.ResponseWriter, r *http.Request, status string) {
    // If the status is not present return an error
    // Later we will add a Response Package and handle all response there, including errors
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

func PostPet(w http.ResponseWriter, r *http.Request, value string) {
    type EmptyResponse struct {
         Code    int    `json:"code"`
        Type    string  `json:"type"`
        Message string  `json:"message"`
        Data    string  `json:"data"`
    }

    var data EmptyResponse
    data.Code = 200
    data.Type = "Success"
    data.Message = "Operation executed successfully"
    data.Data = "PostPet empty response"

    // response as JSON
    response, err := json.Marshal(data)
    if (err != nil) {
        // send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func UpdatePet(w http.ResponseWriter, r *http.Request, ps httprouter.Params, value string) {
    type EmptyResponse struct {
         Code    int    `json:"code"`
        Type    string  `json:"type"`
        Message string  `json:"message"`
        Data    string  `json:"data"`
    }

    var data EmptyResponse
    data.Code = 200
    data.Type = "Success"
    data.Message = "Operation executed successfully"
    data.Data = "UpdatePet empty response"

    // response as JSON
    response, err := json.Marshal(data)
    if (err != nil) {
        // send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func UploadImage(w http.ResponseWriter, r *http.Request, value string) {
    type EmptyResponse struct {
         Code    int    `json:"code"`
        Type    string  `json:"type"`
        Message string  `json:"message"`
        Data    string  `json:"data"`
    }

    var data EmptyResponse
    data.Code = 200
    data.Type = "Success"
    data.Message = "Operation executed successfully"
    data.Data = "UploadImage empty response"

    // response as JSON
    response, err := json.Marshal(data)
    if (err != nil) {
        // send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func DeletePet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    type EmptyResponse struct {
         Code    int    `json:"code"`
        Type    string  `json:"type"`
        Message string  `json:"message"`
        Data    string  `json:"data"`
    }

    var data EmptyResponse
    data.Code = 200
    data.Type = "Success"
    data.Message = "Operation executed successfully"
    data.Data = "DeletePet empty response"

    // response as JSON
    response, err := json.Marshal(data)
    if (err != nil) {
        // send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}
