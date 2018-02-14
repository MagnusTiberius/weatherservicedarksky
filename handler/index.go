package handler

import (
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "unknown"
	}

	response := ResponseMessage{Message: "Hello, " + name, Code: 200}
	JsonResponseWrite(w, response, 200)
}

func DarkSkyHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	lat := q.Get("lat")
	if lat == "" {
		lat = "unknown"
	}
	lng := q.Get("lng")
	if lng == "" {
		lng = "unknown"
	}

	response := ResponseMessage{Message: "lat:" + lat + " lng:" + lng, Code: 200}
	JsonResponseWrite(w, response, 200)
}

func GoogleGeoHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	addr := q.Get("addr")
	if addr == "" {
		addr = "unknown"
	}

	response := ResponseMessage{Message: "addr:" + addr, Code: 200}
	JsonResponseWrite(w, response, 200)
}
