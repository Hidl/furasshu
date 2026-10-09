package main

import(
	"fmt"
	"net/http"
	"database/sql"
	"github.com/hidl/furasshu/internal/database"
	"os"
)

import _ "github.com/lib/pq"

// API config
	type apiConfig struct {
		db *database.Queries
	}

func main()  {

	// Connect to the database
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	} else {
		fmt.Println("Connected to the database successfully!")
	}
	dbQueries := database.New(db)

	// API config
	cfg := apiConfig{
		db: dbQueries,
	}
	
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("POST /addCard", cfg.handleAddCard)

	// Set up server to listen on port 8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	listenErr := server.ListenAndServe()
	if listenErr != nil {
		fmt.Println("Error starting server:", listenErr)
	}

}