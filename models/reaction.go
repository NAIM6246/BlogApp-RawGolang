package models

type Reaction struct {
	ID        int
	PostID    int
	UserID    int
	Is_Like   bool
	Is_Unlike bool
}
