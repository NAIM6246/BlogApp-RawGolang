package repositories

import (
	"fmt"
	"workspacify-blog/db"
	"workspacify-blog/models"
	"workspacify-blog/models/dtos"
)

type ReactionRepository struct {
	db *db.DB
}

func NewReactionRepository() *ReactionRepository {
	return &ReactionRepository{
		db: db.GetDBInstance(),
	}
}

func (repo *ReactionRepository) CreateReaction(data *models.Reaction) (int64, error) {
	res, err := repo.db.Exec(`INSERT INTO REACTIONS 
						(post_id, user_id, is_like, is_unlike)
						VALUES (?, ?, ?, ?)`, data.PostID, data.UserID, data.Is_Like, data.Is_Unlike)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (repo *ReactionRepository) GetPostReactedUser(postID int64, limit int64, lastID int64, liked bool) ([]*dtos.ReactedUserInPost, error) {
	var users []*dtos.ReactedUserInPost

	query := `SELECT P.ID AS POST_ID, U.ID AS USER_ID, U.NAME AS USER_NAME 
			FROM REACTIONS AS R JOIN POSTS AS P ON R.POST_ID = P.ID
			JOIN USERS AS U ON R.USER_ID = U.ID
			WHERE `
	if liked {
		query += `IS_LIKE = TRUE AND P.ID = ?`
	} else {
		query += `IS_UNLIKE = TRUE AND P.ID = ?`
	}
	query += ` AND U.ID > ? LIMIT ?`

	fmt.Println("qu", query)
	res, err := repo.db.Query(query, postID, lastID, limit)
	if err != nil {
		fmt.Println("error ", err)
		return nil, err
	}

	for res.Next() {
		var user dtos.ReactedUserInPost

		err := res.Scan(&user.PostID, &user.UserID, &user.UserName)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
