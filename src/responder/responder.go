package responder

import (
    "encoding/json"
    "net/http"
)

// error response data structures
type Error struct {
    Code    int     `json:"code"`
    Type    string  `json:"type"`
    Message string  `json:"message"`
    Data    struct {
        Error       bool `json:"error"`
        Message     string `json:"message"`
    } `json:"data"`
}

// empty response structure
type Empty struct {
    Code    int    `json:"code"`
    Type    string  `json:"type"`
    Message string  `json:"message"`
    Data    struct {
        Status  bool `json:"status"`
        Message string `json:"message"`
    } `json:"data"`
}

// valid single row response structure - map[string]interface{}
type ResponseSingle struct {
    Code    int     `json:"code"`
    Type    string  `json:"type"`
    Message string  `json:"message"`
    Data    map[string]interface{}
}

// valid multiple row response structure - map[int]interface{}
type ResponseMultiple struct {
    Code    int     `json:"code"`
    Type    string  `json:"type"`
    Message string  `json:"message"`
    Data    map[int]interface{}
}

// SendErrorResponse will ne used to handle error responses
// We can add some logging if we needed to - hence having these responses handled separately
func SendErrorResponse(w http.ResponseWriter, r *http.Request, response Error) {
    res, err := json.Marshal(response)
    if (err != nil) {
        // last chance !! - send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}

// SendEmptyResponse is really a place holder for responding to 'no data found' events - also good for testing etc
func SendEmptyResponse(w http.ResponseWriter, r *http.Request, response Empty) {
    res, err := json.Marshal(response)
    if (err != nil) {
        // last chance !! - send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}

// SendSingleResponse is used to send a single row (single result structures) response
func SendSingleResponse(w http.ResponseWriter, r *http.Request, response ResponseSingle) {
    res, err := json.Marshal(response)
    if (err != nil) {
        // last chance !! - send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}

// SendMultipleResponse is used to return an array (multiple result structure) responses
func SendMultipleResponse(w http.ResponseWriter, r *http.Request, response ResponseMultiple) {
    res, err := json.Marshal(response)
    if (err != nil) {
        // last chance !! - send empty response
        http.NotFound(w, r)
    }

    // set content type to return JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}