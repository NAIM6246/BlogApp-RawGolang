package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"workspacify-blog/configs"
	"workspacify-blog/db"
	"workspacify-blog/handlers"

	"github.com/go-chi/chi/v5"
)

const port = 8005

var initDBTableOnce sync.Once

func init() {
	dbConfig := configs.NewDBConfig()
	dbInstance := db.ConnectDB(dbConfig)

	initDBTableOnce.Do(func() {
		err := dbInstance.Migration()
		if err != nil {
			log.Println("failed to create table ", "error: ", err)
		}

		// err = dbInstance.InsertInitialDataIntoTable()
		// if err != nil {
		// 	log.Println("failed to insert initial batch data", "error: ", err)
		// }
		// err = utils.InitFileWriter()
		// if err != nil {
		// 	log.Println("failed to initialize file writer")
		// }
	})
}

func main() {
	router := chi.NewRouter()

	userHandler := handlers.NewUserHandler()
	postHandler := handlers.NewPostHandler()
	commentHandler := handlers.NewCommentHandler()
	reactionHandler := handlers.NewReactionHandler()

	router.Route("/api", func(r chi.Router) {
		r.Route("/users", userHandler.Handle)
		r.Route("/posts", postHandler.Handle)
		r.Route("/comments", commentHandler.Handle)
		r.Route("/reactions", reactionHandler.Handle)
	})

	// File server
	router.Handle("/files/*", http.StripPrefix("/files/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running on port 8005")
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

// in the example you used os.Create("/uploads") but while you are sending file url you are
