package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Result struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	Geometry          Geometry           `json:"geometry"`
	PlaceID           string             `json:"place_id"`
	Types             []string           `json:"types"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type NeSw struct {
	NorthEast LatLng `json:"northeast"`
	SouthWest LatLng `json:"southwest"`
}

type Geometry struct {
	Bounds       NeSw   `json:"bounds"`
	Location     LatLng `json:"location"`
	LocationType string `json:"location_type"`
	ViewPort     NeSw   `json:"viewport"`
}

type ReturnValue struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/address/{addr}", AddressHandler)
	r.HandleFunc("/address/{addr}/{time}", AddressHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8088", "http://35.226.247.163:8088/"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8090", handler)
}

type ResponseMessage struct {
	Message     ReturnValue `json:"message"`
	Code        int         `json:"code"`
	ReturnValue ReturnValue `json:"returnvalue"`
}

func JsonResponseWrite(w http.ResponseWriter, message interface{}, statusCode int) {

	body, err := json.Marshal(message)

	if statusCode == 200 && err == nil {
		msg, _ := json.Marshal(message)
		w.Header().Set("content-type", "application/json")
		w.Write(msg)
	} else {
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			http.Error(w, string(body), statusCode)
		}
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "unknown"
	}

	response := ResponseMessage{Message: ReturnValue{}, Code: 200}
	JsonResponseWrite(w, response, 200)
}

func AddressHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	addr := vars["addr"]
	if addr == "" {
		addr = "unknown"
		log.Print("addr undefined")
		response := ResponseMessage{Message: ReturnValue{}, Code: 200}
		JsonResponseWrite(w, response, 200)
	}
	tm := vars["time"]
	if tm == "" {
		log.Print("time undefined")
	}

	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + url.QueryEscape(addr) + "&key=AIzaSyC43h_g6FigjhSt3Y46oeqTu-Ydd24F5KI"

	fmt.Printf("url:%s\n\n", url)
	response, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: ReturnValue{}, Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: ReturnValue{}, Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	var rv ReturnValue
	rv = ReturnValue{}

	err = json.Unmarshal(responseData, &rv)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: ReturnValue{}, Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	log.Printf("lat:%f\n\n", rv.Results[0].Geometry.Location.Lat)
	log.Printf("lng:%f\n\n", rv.Results[0].Geometry.Location.Lng)

	dkey := "44881b03349f3f598dbd77b7eaeb215b"
	urlDarkSky := ""
	if tm == "" {
		urlDarkSky = fmt.Sprintf("https://api.darksky.net/forecast/%s/%f,%f", dkey, rv.Results[0].Geometry.Location.Lat, rv.Results[0].Geometry.Location.Lng)
	} else {
		urlDarkSky = fmt.Sprintf("https://api.darksky.net/forecast/%s/%f,%f,%s", dkey, rv.Results[0].Geometry.Location.Lat, rv.Results[0].Geometry.Location.Lng, tm)
	}
	log.Printf("urlDarkSky:%s\n\n", urlDarkSky)

	response, err = http.Get(urlDarkSky)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: ReturnValue{}, Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		response := ResponseMessage{Message: ReturnValue{}, Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseData)
}
