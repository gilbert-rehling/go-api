package main

import (
	"fmt"
	"github.com/gilbert-rehling/go-api/models"
	"github.com/gilbert-rehling/go-api/http"
)

func init() {
    var text string

    text = http.GetPets(db)

    fmt.Println(text)

}

func main() {
    // acknowledge the process has started
    fmt.Println("start")
}
