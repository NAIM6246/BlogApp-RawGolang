package db

func (d *DB) createUserTable() error {
	_, err := d.Exec(`CREATE TABLE IF NOT EXISTS USERS (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		NAME VARCHAR(50) NOT NULL,
		EMAIL VARCHAR(150) NOT NULL,
		PASSWORD VARCHAR(200) NOT NULL)`)
	return err
}

func (d *DB) createCommentTable() error {
	_, err := d.Exec(`CREATE TABLE IF NOT EXISTS COMMENTS (
					ID INT AUTO_INCREMENT PRIMARY KEY, 
		 			DESCRIPTION VARCHAR(500) NOT NULL,
					AUTHOR_ID INT NOT NULL REFERENCES USERS(ID) ON DELETE CASCADE,
					POST_ID INT NOT NULL REFERENCES POSTS(ID) ON DELETE CASCADE)`)
	return err
}

func (d *DB) createPostTable() error {
	_, err := d.Exec(`CREATE TABLE IF NOT EXISTS POSTS (
					ID INT AUTO_INCREMENT PRIMARY KEY, 
					DESCRIPTION VARCHAR(1000) NOT NULL,
					MEDIA VARCHAR(300),
					LIKE_COUNT INT,
					UNLIKE_COUNT INT,
					COMMENT_COUNT INT,
					AUTHOR_ID INT NOT NULL REFERENCES USERS(ID) ON DELETE CASCADE)`)
	return err
}

func (d *DB) createReactionTable() error {
	_, err := d.Exec(`CREATE TABLE IF NOT EXISTS REACTIONS (
		ID INT AUTO_INCREMENT PRIMARY KEY, 
		IS_LIKE BOOL,
		IS_UNLIKE BOOL,
		USER_ID INT NOT NULL REFERENCES USERS(ID) ON DELETE CASCADE,
		POST_ID INT NOT NULL REFERENCES POSTS(ID) ON DELETE CASCADE)`)
	return err
}

func (d *DB) initialInsertIntoDepartmentsTable() error {
	_, err := d.Exec(`INSERT INTO DEPARTMENTS (DEPARTMENT_NAME, DEPARTMENT_CODE) VALUES 
					('English', 'ELA_01'),
					('Mathematics', 'ELA_02'),
					('Bangla', 'ELA_03')`)
	return err
}

func (d *DB) initialInsertIntoTeachersTable() error {
	_, err := d.Exec(`INSERT INTO TEACHERS (TEACHER_NAME, SALARY, DEPARTMENT_ID) VALUES
					('William Shakespeare', 20000, 1),
					('Christopher Marlowe', 15000, 1),
					('John Milton', 12000, 2),
					('John Dryden', 10000, 2),
					('William Wordsworth', 17000, 3),
					('S.T. Coleridge', 11000, 2) `)
	return err
}

func (d *DB) initialInsertIntoStudentsTable() error {
	_, err := d.Exec(`INSERT INTO STUDENTS (STUDENT_NAME, TEACHER_ID) VALUES
					('Stuart Mil', 1),
					('Lord Alfred', 1),
					('Thomas Hardy', 3),
					('Emily Bronte', 3),
					('Leo Tolstoy', 1),
					('Karl Marx', 5)`)
	return err
}
