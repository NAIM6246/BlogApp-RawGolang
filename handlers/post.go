package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"workspacify-blog/models"
	"workspacify-blog/models/dtos"
	"workspacify-blog/repositories"
	"workspacify-blog/utils"

	"github.com/go-chi/chi/v5"
)

type PostHandlers struct {
	postRepository *repositories.PostRepository
}

func NewPostHandler() IHandler {
	return &PostHandlers{
		postRepository: repositories.NewPostRepository(),
	}
}

func (h *PostHandlers) Handle(router chi.Router) {
	router.Post("/", h.createPost)
	router.Get("/", h.getPosts)
}

func (h *PostHandlers) createPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	post.Description = r.FormValue("description")
	post.LikeCount = utils.ParseInt(r.FormValue("like_count"))
	post.UnlikeCount = utils.ParseInt(r.FormValue("unlike_count"))
	post.CommentCount = utils.ParseInt(r.FormValue("comment_count"))
	post.AuthorID = utils.ParseInt(r.FormValue("author_id"))

	filePath, err := h.uploadFile(r)
	if err != nil {
		log.Println("error uploading file", err)
		InternalServerError(w, err)
		return
	}
	log.Println(filePath)

	post.Media = filePath

	id, err := h.postRepository.CreatePost(&post)
	if err != nil {
		log.Println("error creating post", err)
		InternalServerError(w, err)
		return
	}

	Created(w, id)
}

func (h *PostHandlers) getPosts(w http.ResponseWriter, r *http.Request) {
	req := struct {
		LastID int `json:"last_id"`
		Limit  int `json:"limit"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, err)
		return
	}

	posts, err := h.postRepository.GetPosts(req.Limit, req.LastID)
	if err != nil {
		NotFound(w, err)
		return
	}

	Ok(w, &dtos.PaginatedPostResposnse{
		Posts:  posts,
		LastID: int64(posts[len(posts)-1].ID),
	})
}
