package api

import (
	"ManagementSystem/api/generated"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp := []generated.UserDto{
		{
			Id:   uuid.New(),
			Name: "Arbitrary Name",
		},
		{
			Id:   uuid.New(),
			Name: "Another Name",
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
