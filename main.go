package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/vpcraft/feedlygo/internal/db"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("APP_PORT")
	if portString == "" {
		log.Fatal("Environment variable PORT must be set")
		return
	}

	dbUrlString := os.Getenv("DB_URL")
	if dbUrlString == "" {
		log.Fatal("Environment variable DB_URL must be set")
		return
	}

	conn, err := sql.Open("postgres", dbUrlString)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return
	}

	// We can inject this dependency into our handlers to access the database
	db := db.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerReadinessErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	// v1Router.Get("/feeds", apiCfg.middlewareAuth(apiCfg.handlerGetFeedByID))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)
	v1Router.Post("/follows/create", apiCfg.middlewareAuth(apiCfg.handlerFollowToFeed))
	v1Router.Post("/follows/delete", apiCfg.middlewareAuth(apiCfg.handlerUnfollowFromFeed))
	v1Router.Get("/follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v...", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
