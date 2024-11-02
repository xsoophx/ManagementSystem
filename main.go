package main

import (
	"ManagementSystem/api"
	"ManagementSystem/api/generated"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT_HOST")
	dbHost := os.Getenv("DB_HOST_HOST")

	if os.Getenv("DOCKER_ENV") == "true" {
		dbPort = os.Getenv("POSTGRES_PORT_CONTAINER")
		fmt.Println("DB_HOST_CONTAINER:", os.Getenv("DB_HOST_CONTAINER"))
		dbHost = os.Getenv("DB_HOST_CONTAINER")
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	server := api.NewServer(dbUrl)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// TODO: Add auth
	//apiRouter.Use(middlewares.AuthMiddleware)

	generated.HandlerFromMux(server, apiRouter)

	router.HandleFunc("/", homeHandler)

	s := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8123",
	}

	log.Fatal(s.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
