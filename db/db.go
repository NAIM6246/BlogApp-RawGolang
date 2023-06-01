package db

import (
	"fmt"
	"workspacify-blog/models"
)

func GetQuerInterface() *DB {
	return dbInstance
}

func (db *DB) GetAllPost() ([]*models.Post, error) {
	var posts []*models.Post

	res, err := db.Query(`SELECT D.ID AS ID, D.DEPARTMENT_NAME AS DEPARTMENT_NAME, SUM(T.SALARY) as TOTAL_COST
					FROM DEPARTMENTS AS D INNER JOIN TEACHERS AS T
					ON D.ID=T.DEPARTMENT_ID
					GROUP BY D.ID, D.DEPARTMENT_NAME
					ORDER BY TOTAL_COST DESC`)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var post *models.Post
		err := res.Scan(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (db *DB) GetAllComment() ([]*models.Comment, error) {
	var comments []*models.Comment

	res, err := db.Query(`SELECT * FROM COMMENTS`)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var comment *models.Comment
		err := res.Scan(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	fmt.Println(comments)
	return comments, nil
}

// func (db *DB) AddTeacher(data *models.TeacherCreateDto) error {
// 	_, err := db.Exec(`INSERT INTO TEACHERS (TEACHER_NAME, SALARY, DEPARTMENT_ID) VALUES (?, ?, ?)`, data.TeacherName, data.Salary, data.DepartmentID)
// 	return err
// }

// func (db *DB) AddDepartment(data *models.DepartmentCreateDto) error {
// 	_, err := db.Exec(`INSERT INTO DEPARTMENTS (DEPARTMENT_NAME, DEPARTMENT_CODE) VALUES  (?, ?)`, data.DepartmentName, data.DepartmentCode)
// 	return err
// }
