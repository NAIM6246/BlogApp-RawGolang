package models

type Reaction struct {
	ID        int
	PostID    int64
	UserID    int64
	Is_Like   *bool
	Is_Unlike *bool
}
