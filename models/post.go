package models

type Post struct {
	ID           int
	Description  string
	Media        string
	LikeCount    int64
	UnlikeCount  int64
	CommentCount int64
	AuthorID     int64
}
