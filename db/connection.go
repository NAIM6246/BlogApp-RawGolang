package db

import (
	"fmt"
	"log"
	"sync"
	"workspacify-blog/configs"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3" // SQLite driver
)

var (
	connDBOnce sync.Once
	dbInstance *DB
)

type DB struct {
	*sql.DB
}

func connectDB(config *configs.DBConfig) error {
	// need to configure this
	dsn := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/temp")
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection failed")
		return err
	}
	fmt.Println("Database connected successfully.")
	dbInstance = &DB{conn}
	return nil
}

func ConnectDB(config *configs.DBConfig) *DB {
	connDBOnce.Do(func() {
		err := connectDB(config)
		if err != nil {
			panic("failed to connect DB: " + err.Error())
		}
	})
	return dbInstance
}

func GetDBInstance() *DB {
	return dbInstance
}

// Migration at first deletes all the tables if exists and then create new tables
func (db *DB) Migration() error {

	// to avoid duplicate data entry during every run
	// db.deleteTables()

	err := db.createUserTable()
	if err != nil {
		log.Println("error creating users table", "error: ", err)
		return err
	}

	err = db.createPostTable()
	if err != nil {
		log.Println("error creating post table", "error: ", err)
		return err
	}

	err = db.createCommentTable()
	if err != nil {
		log.Println("error creating comment table", "error: ", err)
		return err
	}

	err = db.createReactionTable()
	if err != nil {
		log.Println("error creating reaction table", "error: ", err)
		return err
	}

	return nil
}

// InsertInitialDataIntoTable inserts intial given data into the tables
func (db *DB) InsertInitialDataIntoTable() error {
	err := db.initialInsertIntoDepartmentsTable()
	if err != nil {
		log.Println("error inserting batch data into departments table", "error: ", err)
		return err
	}

	err = db.initialInsertIntoTeachersTable()
	if err != nil {
		log.Println("error inserting batch data into teachers table", "error: ", err)
		return err
	}

	err = db.initialInsertIntoStudentsTable()
	if err != nil {
		log.Println("error inserting batch data into departments table", "error: ", err)
		return err
	}
	return nil
}

// deletes table if exists
func (db *DB) deleteTables() {
	db.Exec(`DROP TABLE COMMENTS`)
	db.Exec(`DROP TABLE REACTIONS`)
	db.Exec(`DROP TABLE POSTS`)
	db.Exec(`DROP TABLE USERS`)
}
