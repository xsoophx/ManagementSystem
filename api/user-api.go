package api

import (
	"ManagementSystem/api/generated"
	"encoding/json"
	"net/http"
)

func (Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp := []generated.User{
		{
			Id:   1,
			Name: "Arbitrary Name",
		},
		{
			Id:   2,
			Name: "Another Name",
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
