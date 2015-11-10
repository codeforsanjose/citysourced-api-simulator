package main

import (
	"CitySourcedAPI/config"
	"CitySourcedAPI/data"
	"CitySourcedAPI/docs"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/request"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	log = logs.Log
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/docs/", docHandler)
	http.HandleFunc("/api/", apiHandler)
	http.ListenAndServe(":5050", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", docs.Home.Title, docs.Home.Body)
}

func docHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/docs/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", docs.Detail.Title, docs.Detail.Body)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.URL.Path != "/api/" || r.Method != "POST" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	// decoder := xml.NewDecoder(r.Body)
	// var stReq request.Request_Type
	// err := decoder.Decode(&stReq)
	// if err != nil {
	// 	fmt.Printf("Error decoding\n")
	// 	errorHandler(w, r, http.StatusNotFound)
	// }

	fmt.Printf("api request - method: %v\n%#v\n", r.Method, string(req))
	resp, err := request.Process(string(req), start)

	fmt.Fprint(w, resp)
}

func init() {
	if err := config.Init("config.json"); err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
	}

	if err := data.Init("data.json"); err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
	}
}
