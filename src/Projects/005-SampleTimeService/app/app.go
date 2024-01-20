package app

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func timeServerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "%s", "Method must be GET!...")
		return
	}
	now := time.Now()
	w.WriteHeader(http.StatusOK)
	ci := ClientInfo{Host: r.Host, Name: "Hello", DateTime: now.Format("02/01/2006 15:04:05")}
	fmt.Fprintf(w, "%v", ci)
}

func Run() {
	if len(os.Args) != 2 {
		fmt.Printf("Wrong number of arguments!...")
		os.Exit(1)
	}

	handler := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + os.Args[1],
		Handler: handler, IdleTimeout: 10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	handler.Handle("/time", http.HandlerFunc(timeServerHandler))

	fmt.Println("Service is waiting for a client")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Message:%s\n", err.Error())
	}
}
