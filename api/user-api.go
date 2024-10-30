package api

import (
	"ManagementSystem/api/generated"
	"encoding/json"
	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"
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

func (Server) GetUsersId(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	resp := generated.UserDto{
		Id:   uuid.New(),
		Name: "Arbitrary Name",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	resp := generated.UserDto{
		Id:   uuid.New(),
		Name: "Arbitrary Name",
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (Server) UpdateUser(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	resp := generated.UserDto{
		Id:   uuid.New(),
		Name: "Arbitrary Name",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
