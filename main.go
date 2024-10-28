package main

import (
	"ManagementSystem/api"
	"ManagementSystem/api/generated"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	server := api.NewServer()

	router := mux.NewRouter()
	handler := generated.HandlerFromMux(server, router)

	router.HandleFunc("/", HomeHandler)

	s := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
