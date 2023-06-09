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

type ReactionHandler struct {
	reactionRepository *repositories.ReactionRepository
	postRepository     *repositories.PostRepository
}

func NewReactionHandler() IHandler {
	return &ReactionHandler{
		reactionRepository: repositories.NewReactionRepository(),
		postRepository:     repositories.NewPostRepository(),
	}
}

func (h *ReactionHandler) Handle(router chi.Router) {
	router.Post("/", h.createReaction)
	router.Get("/reacted-users", h.getReactedUserOfPost)
}

func (h *ReactionHandler) createReaction(w http.ResponseWriter, r *http.Request) {
	var payload dtos.ReactionCreateDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, err)
		return
	}

	id, err := h.reactionRepository.CreateReaction(&models.Reaction{
		PostID:    payload.PostID,
		UserID:    payload.UserID,
		Is_Like:   payload.Is_Like,
		Is_Unlike: payload.Is_Unlike,
	})
	if err != nil {
		InternalServerError(w, err)
		return
	}
	err = h.postRepository.UpdatePostCount(payload.PostID, *payload.Is_Like, false)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	log.Println("reaction created", "data", payload)
	Created(w, id)
}

func (h *ReactionHandler) getReactedUserOfPost(w http.ResponseWriter, r *http.Request) {
	var payload dtos.GetReactedUserOfPostReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, err)
		return
	}

	users, err := h.reactionRepository.GetPostReactedUser(payload.PostID, payload.Limt, payload.LastID, payload.Liked)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	var lastID int64
	if len(users) > 0 {
		lastID = int64(users[len(users)-1].UserID)
	}

	log.Println("users fetched", "users", users)
	Created(w, &dtos.PaginatedReactedUSsers{
		Users:  users,
		LastID: lastID,
	})
}
