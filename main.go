package main

import (
	"ManagementSystem/api"
	"ManagementSystem/api/generated"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbUrl := "postgres://" + dbUser + ":" + dbPassword + "@localhost:5432/" + dbName
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
