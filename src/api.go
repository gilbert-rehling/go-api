package main

// this router has issues !!  "github.com/julienschmidt/httprouter"
import (
    "fmt"
    "time"
	"log"
	"os"
	"net/http"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"github.com/gilbert-rehling/go-api/db"
    "github.com/gilbert-rehling/go-api/handlers"
    "gopkg.in/yaml.v2"
)

// define the config structure
type Config struct {
    Server struct {
        Host string `yaml:"host"`
        Port string `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        Type    string `yaml:"type"`
        Host    string `yaml:"host"`
        Port    string `yaml:"port"`
        User    string `yaml:"username"`
        Pwd     string `yaml:"password"`
        Name    string `yaml:"name"`
    } `yaml:"database"`
}
var config Config

// initialises the application
func init() {
    // init started

    // populate the config from file
    readFile(&config)

    // view the config data (dev)
    // fmt.Printf("%+v", config)  // print the config contents

    // init finished
}

// process any errors
func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}

// process configuration from file
func readFile(config *Config) {

    // get the current directory
    dir, err := os.Getwd()
    if err != nil {
        processError(err)
    }
    fmt.Println(dir) // print the current directory
    var configPath = dir + "/config/db.yaml"
    fmt.Println(configPath)  // print the path to config

    f, err := os.Open(configPath)
    if err != nil {
        processError(err)
    }
    defer f.Close()

    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(config)
    if err != nil {
        processError(err)
    }
}

// load the main app
func main() {
    // Set up the DB connection
    // Handle the incoming route requests
    // Load the listener | server
    var err error

    // create the DB link !! todo: using hard coded credentials for testing !!
    db.Conn, err = sql.Open(config.Database.Type, config.Database.User + ":" + config.Database.Pwd + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Name)
    if err != nil {
        log.Panic("DB connection error: %s", err)
    }

    // To avoid error due to connection being closed on the server
    db.Conn.SetConnMaxLifetime(14000 * time.Second)

    // check that the db link works
    err = db.Conn.Ping()
    if err != nil {
        log.Panic("DB ping error: %s", err)
    }

    // Using imported router library
    // !! this flashy router wont work correctly for me at the moment
    router := httprouter.New()

    // This router isn't working for me !! keep getting 404's
    // But this 'wildcard' route works !!
    router.GET("/pet/*segment", handlers.GetSwitch)

    // Send to the POST switch
    router.POST("/pet/*segment", handlers.PostSwitch)

    // There is only 1 DELETE route
    router.DELETE("/pet/*segment", handlers.DeletePet)

    // There is only 1 PUT route
    // This needs to share the UpdatePet handler
    router.PUT("/pet/*segment", handlers.PutSwitch)

    // set to listen on external facing IP address on DEV box
    log.Fatal(http.ListenAndServe(config.Server.Host + ":" + config.Server.Port, router))
}
