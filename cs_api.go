package main

import (
    "CitySourcedAPI/request"
    
	"fmt"
    "io/ioutil"
	"net/http"
)

var docs_home = `
This is the documentation for the City Sourced test API System.
Enter URL "docs" for more information.
`
var docs_detail = `
This is detailed documentation for the City Sourced test API System.

Path 
`

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
	fmt.Fprint(w, docs_home)
}

func docHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/docs/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, docs_detail)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
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
    request.Process(string(req))
    
	fmt.Fprint(w, "call to api!")
}
