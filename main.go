package main

import (
	"ManagementSystem/api"
	"ManagementSystem/api/generated"
	"ManagementSystem/util"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := createLogger()
	ctx := context.WithValue(context.Background(), "logger", logger)

	dbUrl := createDbUrl(ctx)
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

func createDbUrl(ctx context.Context) string {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT_HOST")
	dbHost := os.Getenv("DB_HOST_HOST")

	if os.Getenv("DOCKER_ENV") == "true" {
		dbPort = os.Getenv("POSTGRES_PORT_CONTAINER")
		dbHost = os.Getenv("DB_HOST_CONTAINER")
	}

	logger := util.GetLogger(ctx)
	logger.Info("Env values for Database connection",
		zap.String("POSTGRES_USER", dbUser),
		zap.String("POSTGRES_DB", dbName),
		zap.String("POSTGRES_PORT_HOST", dbPort),
		zap.String("DB_HOST_HOST", dbHost),
	)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func createLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger
}
