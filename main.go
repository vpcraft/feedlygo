package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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

	fmt.Println("Port:", portString)
	fmt.Println("DB_URL:", dbUrlString)

	conn, err := sql.Open("postgres", dbUrlString)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return
	}

	apiCfg := apiConfig{
		DB: db.New(conn),
	}

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
	// v1Router.Mount()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerReadinessErr)

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
