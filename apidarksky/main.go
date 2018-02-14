package main

import (
	"net/http"

	"github.com/MagnusTiberius/weatherservicedarksky/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.DarkSkyHandler)
	r.HandleFunc("/weather/{lat}/{lng}", handler.DarkSkyHandler)
	http.ListenAndServe(":8090", r)
}
