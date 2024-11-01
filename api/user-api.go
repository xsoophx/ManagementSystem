package api

import (
	"ManagementSystem/api/generated"
	"ManagementSystem/services/models"
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
	var userDto generated.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := s.UserService.RegisterUser(userDto.FirstName, userDto.LastName, userDto.Email)

	if err != nil {
		http.Error(w, "Could not register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(userToUserDto(user))
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

func userToUserDto(user *models.User) *generated.UserDto {
	return &generated.UserDto{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
