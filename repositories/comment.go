package repositories

import (
	"fmt"
	"workspacify-blog/db"
	"workspacify-blog/models"
)

type CommentRepository struct {
	db *db.DB
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		db: db.GetDBInstance(),
	}
}

func (repo *CommentRepository) CreateComment(data *models.Comment) (int64, error) {
	res, err := repo.db.Exec(`INSERT INTO COMMENTS 
							(description , post_id, author_id)
							VALUES (?, ?, ?)`, data.Description, data.PostID, data.AuthorID)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (repo *CommentRepository) GetPostComments(postID int, limit int, lastID int) ([]*models.Comment, error) {
	var comments []*models.Comment

	res, err := repo.db.Query(`SELECT * FROM COMMENTS WHERE ID > ? AND POST_ID = ? LIMIT ?`, lastID, postID, limit)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var comment models.Comment
		err := res.Scan(&comment.ID, &comment.Description, &comment.AuthorID, &comment.PostID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	fmt.Println(comments)
	return comments, nil
}
