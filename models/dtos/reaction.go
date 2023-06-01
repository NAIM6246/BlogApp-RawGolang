package dtos

type ReactionCreateDto struct {
	PostID    int64 `json:"post_id"`
	UserID    int64 `json:"user_id"`
	Is_Like   *bool `json:"is_like"`
	Is_Unlike *bool `json:"is_unlike"`
}

type ReactedUserInPost struct {
	PostID   int64  `json:"post_id"`
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type GetReactedUserOfPostReq struct {
	PostID int64 `json:"post_id"`
	Limt   int64 `json:"limit"`
	LastID int64 `json:"last_id"`
	Liked  bool  `json:"liked"`
}
