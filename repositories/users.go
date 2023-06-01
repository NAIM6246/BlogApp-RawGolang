package repositories

import (
	"workspacify-blog/db"
	"workspacify-blog/models"
)

type UserRepository struct {
	db *db.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDBInstance(),
	}
}

func (repo *UserRepository) GetAllUser() ([]*models.User, error) {
	var users []*models.User

	res, err := repo.db.Query(`SELECT * FROM USERS`)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var user models.User
		err := res.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (repo *UserRepository) AddeUser(data *models.User) (int64, error) {
	res, err := repo.db.Exec(`INSERT INTO USERS (NAME, EMAIL, PASSWORD) VALUES (?, ?, ?)`, data.Name, data.Email, data.Password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
