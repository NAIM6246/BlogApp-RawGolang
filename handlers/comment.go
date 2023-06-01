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

type CommentHandler struct {
	commentRepository *repositories.CommentRepository
}

func NewCommentHandler() IHandler {
	return &CommentHandler{
		commentRepository: repositories.NewCommentRepository(),
	}
}

func (h *CommentHandler) Handle(router chi.Router) {
	router.Post("/", h.createComment)
	router.Get("/post-comments", h.getPostComments)
}

func (h *CommentHandler) createComment(w http.ResponseWriter, r *http.Request) {
	var payload dtos.CommentCreateDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, err)
		return
	}

	id, err := h.commentRepository.CreateComment(&models.Comment{
		Description: payload.Description,
		PostID:      payload.PostID,
		AuthorID:    payload.AuthorID,
	})
	if err != nil {
		InternalServerError(w, err)
		return
	}

	log.Println("comment created", "data", payload)
	Created(w, id)
}

func (h *CommentHandler) getPostComments(w http.ResponseWriter, r *http.Request) {
	req := struct {
		PostID int `json:"post_id"`
		LastID int `json:"last_id"`
		Limit  int `json:"limit"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, err)
		return
	}

	comments, err := h.commentRepository.GetPostComments(req.PostID, req.Limit, req.LastID)
	if err != nil {
		NotFound(w, err)
		return
	}

	Ok(w, &dtos.PaginatedCommentResponse{
		Comments: comments,
		LastID:   int64(comments[len(comments)-1].ID),
	})
}
