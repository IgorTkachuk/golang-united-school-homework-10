package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", paramHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", badHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", dataPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", headersPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/", rootHandler).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(notFound)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func rootHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func badHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func paramHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	param := mux.Vars(r)["PARAM"]
	w.Write([]byte("Hello, " + param + "!"))
}

func dataPostHandler (w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(append([]byte("I got message:\n"), b...))
}

func headersPostHandler (w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))

	w.Header().Set("a+b", strconv.Itoa(a + b))
	w.WriteHeader(http.StatusOK)
}

func notFound (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
