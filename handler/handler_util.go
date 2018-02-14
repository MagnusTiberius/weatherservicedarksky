package handler

import (
	"encoding/json"
	"net/http"
)

func JsonResponseWrite(w http.ResponseWriter, message interface{}, statusCode int) {

	body, err := json.Marshal(message)

	if statusCode == 200 && err == nil {
		body, _ := json.Marshal(message)
		w.Header().Set("content-type", "application/json")
		w.Write(body)
	} else {
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			http.Error(w, string(body), statusCode)
		}
	}
}
