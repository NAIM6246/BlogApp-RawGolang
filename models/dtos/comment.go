package dtos

import "workspacify-blog/models"

type CommentCreateDto struct {
	Description string `json:"description"`
	PostID      int64  `json:"post_id"`
	AuthorID    int64  `json:"author_id"`
}

type PaginatedCommentResponse struct {
	Comments []*models.Comment
	LastID   int64
}
