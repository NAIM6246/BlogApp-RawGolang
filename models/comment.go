package models

type Comment struct {
	ID          int
	Description string
	AuthorID    int
	PostID      int
}
