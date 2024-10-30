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
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// TODO: Add auth
	//apiRouter.Use(middlewares.AuthMiddleware)

	generated.HandlerFromMux(server, apiRouter)

	router.HandleFunc("/", HomeHandler)

	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
