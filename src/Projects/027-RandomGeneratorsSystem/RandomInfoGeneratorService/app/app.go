package app

import (
	"SampleTimeServiceApp/app/jsondata"
	"SampleTimeServiceApp/csd/util/str"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func checkMethod(w http.ResponseWriter, r *http.Request, message string) bool {
	if r.Method != http.MethodGet {
		http.Error(w, message, http.StatusForbidden)
		return false
	}
	return true
}

func checkError(err error, w http.ResponseWriter, message string, statusCode int) bool {
	if err != nil {
		http.Error(w, message, statusCode)
		return false
	}

	return true
}

func checkParameter(w http.ResponseWriter, r *http.Request, parameter, message string, statusCode int) (bool, string) {
	value := r.FormValue(parameter)
	if value == "" {
		http.Error(w, message, statusCode)
		return false, ""
	}

	return true, value
}

func randomInfoHandler(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, "Method must be GET!...") {
		return
	}

	w.WriteHeader(http.StatusOK)
	n := rand.Intn(20001) + 10000
	min := rand.Intn(6) + 5
	max := rand.Intn(20-min+1) + min
	count := rand.Intn(20) + 10
	basePath := str.GenerateRandomTextEN(10)
	info := jsondata.NewInfo(n, min, max, count, basePath)
	data, err := json.Marshal(info)
	if !checkError(err, w, "Internal server error!...", http.StatusInternalServerError) {
		return
	}
	_, err = fmt.Fprintf(w, "%s", string(data))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Error(w, "No default url", http.StatusNotAcceptable)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func checkArguments() {
	if len(os.Args) != 2 {
		fmt.Printf("Wrong number of arguments!...")
		os.Exit(1)
	}
}

func createServerInfo(portInfo string) (*http.ServeMux, *http.Server) {
	handler := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + portInfo,
		Handler: handler, IdleTimeout: 10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	return handler, server
}

func addHandlers(handler *http.ServeMux) {
	handler.Handle("/info/random", http.HandlerFunc(randomInfoHandler))
	handler.Handle("/", http.HandlerFunc(defaultHandler))
}

func startServer(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Message:%s\n", err.Error())
	}
}

func Run() {
	checkArguments()
	handler, server := createServerInfo(os.Args[1])
	addHandlers(handler)
	fmt.Println("Service is waiting for a client")
	startServer(server)
}
