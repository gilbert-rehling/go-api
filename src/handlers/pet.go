package handlers

// imports saved for later "log" "github.com/julienschmidt/httprouter"   "fmt"
import (
    "strings"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gilbert-rehling/go-api/models"
	"github.com/gilbert-rehling/go-api/responder"
)

type Data struct {
    Error   bool    `json:"error"`
    Message string  `json:"message"`
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
    if (len(id) == 0) {
        // send error response
        var e responder.Error
        e.Code = 200
        e.Type = "Success"
        e.Message = "Operation executed successfully"
        e.Data.Error = true
        e.Data.Message = "GetPet: id missing"
        responder.SendErrorResponse(w, r, e)

    } else {
         // run query on model
         result := models.FindPetById( id )

         if (len(result) == 0) {
            // empty result
            var empty responder.Empty
            empty.Code = 200
            empty.Type = "Success"
            empty.Message = "Operation executed successfully"
            empty.Data.Status = false
            empty.Data.Message = "GetPet: empty response for ID " + id

            responder.SendEmptyResponse(w, r, empty)

         } else {
            var rtn responder.ResponseSingle
            rtn.Code = 200
            rtn.Type = "Success"
            rtn.Message = "Operation executed successfully"
            rtn.Data = result

            responder.SendSingleResponse(w, r, rtn)
         }
    }
}

//  GetPetsByStatus calls FindPetsByStatus
func GetPetsByStatus(w http.ResponseWriter, r *http.Request, status string) {
    // If the status is not present return an error
    // Later we will add a Response Package and handle all response there, including errors
    if (len(status) <= 1) {
        // send error response
        var e responder.Error
        e.Code = 200
        e.Type = "Success"
        e.Message = "Operation executed successfully"
        e.Data.Error = true
        e.Data.Message = "FindPetsByStatus: status missing"

        responder.SendErrorResponse(w, r, e)

    } else {
        // run query on model
        results := models.FindPetsByStatus(status)

        if (len(results) == 0) {
            // empty result
            var empty responder.Empty
            empty.Code = 200
            empty.Type = "Success"
            empty.Message = "Operation executed successfully"
            empty.Data.Status = false
            empty.Data.Message = "FindPetsByStatus: empty response for Status " + status

            responder.SendEmptyResponse(w, r, empty)

        } else {
            // Success response structure
            var rtn responder.ResponseMultiple
            rtn.Code = 200
            rtn.Type = "Success"
            rtn.Message = "Operation executed successfully"
            rtn.Data = results

            responder.SendMultipleResponse(w, r, rtn)
        }
    }
}

func PostPet(w http.ResponseWriter, r *http.Request, value string) {
    // empty response
    var empty responder.Empty
    empty.Code = 200
    empty.Type = "Success"
    empty.Message = "Operation executed successfully"
    empty.Data.Status = false
    empty.Data.Message = "PostPet: empty response"

    responder.SendEmptyResponse(w, r, empty)
}

func UpdatePet(w http.ResponseWriter, r *http.Request, ps httprouter.Params, value string) {
    // empty response
    var empty responder.Empty
    empty.Code = 200
    empty.Type = "Success"
    empty.Message = "Operation executed successfully"
    empty.Data.Status = false
    empty.Data.Message = "UpdatePet: empty response"

    responder.SendEmptyResponse(w, r, empty)
}

func UploadImage(w http.ResponseWriter, r *http.Request, value string) {
    // empty response
    var empty responder.Empty
    empty.Code = 200
    empty.Type = "Success"
    empty.Message = "Operation executed successfully"
    empty.Data.Status = false
    empty.Data.Message = "UploadImage: empty response"

    responder.SendEmptyResponse(w, r, empty)
}

func DeletePet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // empty response
    var empty responder.Empty
    empty.Code = 200
    empty.Type = "Success"
    empty.Message = "Operation executed successfully"
    empty.Data.Status = false
    empty.Data.Message = "DeletePet: empty response"

    responder.SendEmptyResponse(w, r, empty)
}
