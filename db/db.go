package db

func GetQuerInterface() *DB {
	return dbInstance
}

// func (db *DB) AddTeacher(data *models.TeacherCreateDto) error {
// 	_, err := db.Exec(`INSERT INTO TEACHERS (TEACHER_NAME, SALARY, DEPARTMENT_ID) VALUES (?, ?, ?)`, data.TeacherName, data.Salary, data.DepartmentID)
// 	return err
// }

// func (db *DB) AddDepartment(data *models.DepartmentCreateDto) error {
// 	_, err := db.Exec(`INSERT INTO DEPARTMENTS (DEPARTMENT_NAME, DEPARTMENT_CODE) VALUES  (?, ?)`, data.DepartmentName, data.DepartmentCode)
// 	return err
// }
