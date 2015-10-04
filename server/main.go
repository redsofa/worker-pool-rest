package main

import (
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	domain "github.com/redsofa/worker-pool-rest/domain"
  	"io/ioutil"
  	"fmt"
)

//CalcsHandler is a struct that implements the http.Handler interface
type CalcsHandler struct {}

// ServeHTTP method is bound to the CalcsHandler struct. 
// Implementing the http.Handler interface
func (handler *CalcsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// For more information on making Making a RESTful JSON API's in Go see:
	// - http://thenewstack.io/make-a-restful-json-api-go/

	//The io.LimitReader function is to protect again malicious attacks on the server
	//Prevents from having to deal with very large files...

	// LimitReader returns a Reader that reads from r but stops with EOF after
    // n bytes. We limit this to 2 MB        
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 2048666))
    if err != nil {
        log.Println(err)
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        log.Println(err)
        panic(err)
    }

   	//The input list that the user submitted to the server
   	//The ProcessJsonInput function returns a map of Input stucts
    inputCollection := domain.ProcessJsonInput(body)
    //Make sure the user submitted something...
    if len(inputCollection) <= 0 {
        http.Error(w, "{\"error\": \"empty input list submitted.\"}", 500)
        return
    }

 	//The res map will be used to collect results of our workers...
    res := make(map[int]domain.Output)

	//For now we'll just fake the process that does the concurrent calculations
    //and return 0s for the results
    for _, v := range inputCollection {
       output := domain.Output{v.Index, v.NumA, v.NumB, 0}

       res[v.Index] = output
    }

    //Set response header information
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)

	//Create json with all results...
    response, err := domain.GenerateJsonOutput(res)
    if err != nil {
    	log.Println(err)
        panic(err)
    }

    fmt.Fprintf(w, string(response))
}


func main() {
    //The port our server will be litening on
    port := 8080 

	//Our fancy Gorilla mux router
	router := mux.NewRouter() 

	//Our Routes : 
	//We're only interested in POST requests on /calcs URL
	router.Handle("/calcs", &CalcsHandler{}).Methods("POST")  


	log.Printf("Starting server. Listening on port %d", port)
	err := http.ListenAndServe(":" + strconv.Itoa(port), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}