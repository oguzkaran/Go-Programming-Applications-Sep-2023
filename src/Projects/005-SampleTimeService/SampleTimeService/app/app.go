/*
	Sınıf çalışması: Aşağıdaki GET servisleri ekleyiniz
	- ~/date/birth?name=Oguz&surname=Karan&day=10&dmonth=9&year=1976: Verilen bilgilere göre aşağıdaki json formatına dönen
	servis
		{"fullName": "Oğuz Karan, "birthDate": "10/09/1976", "age":47.39, "today":"20/01/2024"}

	Birth servisindeki bilgileri bir veritabanına ekleyiniz
*/

package app

import (
	"SampleTimeServiceApp/app/jsondata"
	"encoding/json"
	"fmt"
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

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, "Method must be GET!...") {
		return
	}

	result, name := checkParameter(w, r, "name", `Parameter "name" required!...`, http.StatusBadRequest)
	if !result {
		return
	}
	now := time.Now()
	w.WriteHeader(http.StatusOK)
	ci := jsondata.NewClientInfo(r.RemoteAddr, "Hello "+name, now.Format("02/01/2006 15:04:05"))
	data, err := json.Marshal(ci)
	if !checkError(err, w, "Internal server error!...", http.StatusInternalServerError) {
		return
	}
	_, err = fmt.Fprintf(w, "%s", string(data))

	if !checkError(err, w, "Internal server error!...", http.StatusInternalServerError) {
		return
	}
}

func dateHandler(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, "Method must be GET!...") {
		return
	}

	result, name := checkParameter(w, r, "name", `Parameter "name" required!...`, http.StatusBadRequest)
	if !result {
		return
	}
	now := time.Now()
	w.WriteHeader(http.StatusOK)
	ci := jsondata.NewClientInfo(r.RemoteAddr, "Hello "+name, now.Format("02/01/2006"))
	data, err := json.Marshal(ci)
	if !checkError(err, w, "Internal server error!...", http.StatusInternalServerError) {
		return
	}
	_, err = fmt.Fprintf(w, "%s", string(data))

	if !checkError(err, w, "Internal server error!...", http.StatusInternalServerError) {
		return
	}
}

func birthDateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under construction")
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
	handler.Handle("/time", http.HandlerFunc(timeHandler))
	handler.Handle("/date", http.HandlerFunc(dateHandler))
	handler.Handle("/date/birth", http.HandlerFunc(birthDateHandler))
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
