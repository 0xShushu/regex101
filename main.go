package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/0xshushu/regex101/handlers"
)

func main() {
	//make a new router
	r := mux.NewRouter().StrictSlash(true)

	//set routes
	r.HandleFunc("/", handlers.Index).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	//start serving
	fmt.Println(http.ListenAndServe(":3000", r))
}