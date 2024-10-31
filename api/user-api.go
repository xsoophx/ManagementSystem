package api

import (
	"ManagementSystem/api/generated"
	"encoding/json"
	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"net/http"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp := []generated.UserDto{
		{
			Id:        uuid.New(),
			FirstName: "Arbitrary First Name",
			LastName:  "Arbitrary Last Name",
		},
		{
			Id:        uuid.New(),
			FirstName: "Another First Name",
			LastName:  "Another Last Name",
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetUsersId(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	user, err := s.UserService.GetUser(id)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	resp := generated.UserDto{
		Id:        uuid.New(),
		FirstName: "Arbitrary First Name",
		LastName:  "Arbitrary Last Name",
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	resp := generated.UserDto{
		Id:        uuid.New(),
		FirstName: "Arbitrary First Name",
		LastName:  "Arbitrary Last Name",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
