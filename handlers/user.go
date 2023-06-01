package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"workspacify-blog/models"
	"workspacify-blog/models/dtos"
	"workspacify-blog/repositories"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

func NewUserHandler() IHandler {
	return &UserHandler{
		userRepository: repositories.NewUserRepository(),
	}
}

func (h *UserHandler) Handle(router chi.Router) {
	router.Get("/", h.getAllUser)
	router.Post("/", h.createUser)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var payload dtos.UserCreateDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, err)
		return
	}

	id, err := h.userRepository.AddeUser(&models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		InternalServerError(w, err)
		return
	}

	log.Println("user created", "data", payload)
	Created(w, id)
}

func (h *UserHandler) getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepository.GetAllUser()
	if err != nil {
		NotFound(w, err)
		return
	}

	Ok(w, users)
}
