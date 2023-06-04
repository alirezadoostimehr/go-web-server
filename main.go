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

func handleName1(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["name"][0]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello " + name))
}

func handleName2(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	w.Write(append([]byte("hello "), name...))
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/bye", handleBye)
	http.HandleFunc("/name1", handleName1)
	http.HandleFunc("/name2", handleName2)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		panic(err)
	}
}
