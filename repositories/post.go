package repositories

import (
	"workspacify-blog/db"
	"workspacify-blog/models"
)

type PostRepository struct {
	db *db.DB
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		db: db.GetDBInstance(),
	}
}

func (repo *PostRepository) GetPosts(limit int, lastID int) ([]*models.Post, error) {
	var posts []*models.Post

	res, err := repo.db.Query(`SELECT * FROM POSTS WHERE ID > ? LIMIT ?`, lastID, limit)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var post models.Post

		err := res.Scan(&post.ID, &post.Description, &post.Media, &post.LikeCount, &post.UnlikeCount, &post.CommentCount, &post.AuthorID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (repo *PostRepository) CreatePost(data *models.Post) (int64, error) {
	res, err := repo.db.Exec(`INSERT INTO POSTS 
							(description , media, like_count, unlike_count, comment_count, author_id)
							 VALUES (?, ?, ?, ?, ?, ?)`,
		data.Description, data.Media, data.LikeCount,
		data.UnlikeCount, data.CommentCount, data.AuthorID)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
