package main

import (
	"CitySourcedAPI/config"
	"CitySourcedAPI/data"
	"CitySourcedAPI/docs"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/request"
	"flag"
	"os"
	"os/signal"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	log   = logs.Log
	Debug = false
)

func main() {
	// setOptions()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/docs/", docHandler)
	http.HandleFunc("/api/", apiHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port()), nil)
}

func init() {
	var port int64
	flag.BoolVar(&Debug, "debug", false, "Activate debug prints to the console.  Active if either this or the value in 'config.json' are set.")
	flag.Int64Var(&port, "port", 0, "Port.  If set, this flag will override the value in 'config.json'.")
	flag.Parse()

	logs.Init(Debug)

	if err := config.Init("config.json", port); err != nil {
		log.Error("Error loading config file: %s\n", err)
	}

	if err := data.Init("data.json"); err != nil {
		log.Error("Error loading config file: %s\n", err)
	}
	log.Info("Running on port: %d", config.Port())

	go SignalHandler(make(chan os.Signal, 1))
	fmt.Println("Press Ctrl-C to shutdown...")
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
	// 	log.Error("Error decoding\n")
	// 	errorHandler(w, r, http.StatusNotFound)
	// }

	log.Debug("api request - method: %v\n%#v\n", r.Method, string(req))
	resp, err := request.Process(string(req), start)

	fmt.Fprint(w, resp)
}

func SignalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("Ctrl-C Received!")
			stop()
			os.Exit(0)
		case os.Kill:
			fmt.Println("SIGKILL Received!")
			stop()
			os.Exit(1)
		}
	}
}

func stop() error {
	return nil
}
