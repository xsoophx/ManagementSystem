package api

import (
	api "ManagementSystem/api/generated"
	"encoding/json"
	"net/http"
)

func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := api.Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
