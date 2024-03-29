package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cdrojasm/golang_freeCodeCamp/project/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // this package is installed even if this is not called in my code
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Loading Secrets...")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("port is not found in the environment")
	}
	dbURLString := os.Getenv("DB_URL")
	if dbURLString == "" {
		log.Fatal("db url is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURLString)
	if err != nil {
		log.Fatal("Cant connect to DB", err)
	}
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScrapping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feed", apiCfg.handleGetFeeds)
	v1Router.Post("/feedFollow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feedFollow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollowsByID))
	v1Router.Delete("/feedFollow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))
	v1Router.Get("/post", apiCfg.middlewareAuth(apiCfg.handlerGetUserPosts))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v ", portString)

	err = srv.ListenAndServe()
	if err == nil {
		log.Fatal(err)
	}

}
