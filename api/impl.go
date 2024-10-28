package api

import (
	api "ManagementSystem/api/generated"
	"encoding/json"
	"net/http"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := api.Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
