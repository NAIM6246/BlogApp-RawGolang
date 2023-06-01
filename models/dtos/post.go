package dtos

import "workspacify-blog/models"

type PostCreateDto struct {
	Description  string `json:"description"`
	LikeCount    int64  `json:"like_count"`
	UnlikeCount  int64  `json:"unlike_count"`
	CommentCount int64  `json:"comment_count"`
	AuthorID     int64  `json:"author_id"`
}

type PaginatedPostResposnse struct {
	Posts  []*models.Post
	LastID int64
}
