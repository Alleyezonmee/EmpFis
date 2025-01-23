package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/Alleyezonmee/EmpFis/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("SERVER_PORT")

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

	// ------ STARTING SERVER ---------
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting at %v", portString)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
