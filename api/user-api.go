package api

import (
	"ManagementSystem/api/generated"
	"ManagementSystem/services/models"
	"ManagementSystem/util"
	"encoding/json"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"go.uber.org/zap"
	"net/http"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	logger := util.GetLogger(s.ctx)
	users, err := s.UserService.GetAllUsers()
	if err != nil {
		msg := "Could not get users"
		logger.Error(msg, zap.Error(err))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	response := make([]generated.UserDto, len(users))
	for i, user := range users {
		response[i] = *userToUserDto(&user)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (s *Server) GetUsersId(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	logger := util.GetLogger(s.ctx)
	user, err := s.UserService.GetUser(id)
	if err != nil {
		msg := "Could not get user"
		logger.Error(msg, zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userToUserDto(user))
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto generated.CreateUserDto
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
	logger := util.GetLogger(s.ctx)
	var userDto generated.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		msg := "Invalid request payload"
		logger.Error(msg, zap.Error(err))
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	user, err := s.UserService.UpdateUser(id, userDto.FirstName, userDto.LastName, userDto.Email)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userToUserDto(user))
}

func userToUserDto(user *models.User) *generated.UserDto {
	return &generated.UserDto{
		Id:        user.ID,
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
		Email:     &user.Email,
		CreatedAt: *user.CreatedAt,
	}
}
