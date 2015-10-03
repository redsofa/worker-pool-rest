package main

import (
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

//CalcsHandler is a struct that implements the http.Handler interface
type CalcsHandler struct {}

//ServeHTTP method is bound to the CalcsHandler struct. 
//Implementing the Handler interface
func (handler *CalcsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I will perform calculations for you...")
}


func main() {
    //The port our server will be litening on
    port := 8080 

	//Our fancy Gorilla mux router
	router := mux.NewRouter() 

	//Our Routes : 
	//For now, will route all (GET,POST,etc) requests on /calcs URL to 
	// CalcsHandler's ServeHTTP function
	router.Handle("/calcs", &CalcsHandler{}).Methods("POST")  


	log.Printf("Starting server. Listening on port %d", port)
	err := http.ListenAndServe(":" + strconv.Itoa(port), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}