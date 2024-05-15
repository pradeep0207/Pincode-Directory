package main

import (
	"fmt"
    "log"  
    "net/http"
	// "example/Go-Api-tutorial/routes"
	"github.com/gorilla/mux"

	controllers "example/Go-Api-tutorial/controllers"
)

const portNum string = ":8080"

func main() {

    fmt.Println("Started on port", portNum)

	router := mux.NewRouter()
    router.HandleFunc("/deletedata", controllers.DeletePersonEndpoint).Methods("DELETE")
	// router.HandleFunc("/create", controllers.CreateData).Methods("POST")
	// router.HandleFunc("/get", controllers.GetData).Methods("GET")

    // Use the router as the handler for all requests
    http.Handle("/", router)

    err := http.ListenAndServe(portNum, nil)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("To close connection CTRL+C :-)")
}