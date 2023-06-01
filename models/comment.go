package models

type Comment struct {
	ID          int64
	Description string
	AuthorID    int64
	PostID      int64
}
