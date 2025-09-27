package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"fem.com/movie-site/data"
	"fem.com/movie-site/handlers"
	"fem.com/movie-site/logger"

	"github.com/joho/godotenv"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Failed to initialise logger $v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	// Log Initializer
	logInstance := initializeLogger()

	movieHandler := handlers.MovieHandler{}

	// Environmental Variables
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// Databse connection
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	defer db.Close()

	// Initialize Data Repository for Movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize Repository")
	}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	// Handler for static files (frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Serving the files")

	const addr = ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed", err)
		log.Fatalf("Server failed: %v", err)
	}
}
