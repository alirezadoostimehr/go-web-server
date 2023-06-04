package main

import (
	"net/http"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}

func handleBye(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
	w.Write([]byte("bye"))
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/bye", handleBye)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		panic(err)
	}
}
