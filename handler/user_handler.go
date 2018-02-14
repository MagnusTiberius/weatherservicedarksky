package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["username"]

	JsonResponseWrite(w, map[string]string{"message": "Hello, " + name}, 200)
}
