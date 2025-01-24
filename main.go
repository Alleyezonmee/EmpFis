package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	handlers "github.com/Alleyezonmee/EmpFis/handlers"
	"github.com/Alleyezonmee/EmpFis/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("SERVER_PORT")
	if portString == "" {
		log.Fatal("Port string is not found in environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DBUrl is not found in environment")
	}

	conn, dbErr := sql.Open("mysql", dbUrl)
	if dbErr != nil {
		log.Fatal("Failed to connect DB")
	}

	apiCfg := ApiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// ------- SETTING ROUTES AND HANDLERS --------
	router.Get("/healthZ", handlers.HandlerReadiness)
	router.Post("/addEmp", apiCfg.HandlerCreateUser)

	// ------ STARTING SERVER ---------
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting at %v", portString)

	serverErr := srv.ListenAndServe()

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
