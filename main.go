package main

import (
	"fmt"
	"os"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)	

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}
	router := chi.NewRouter();
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	srv := &http.Server {
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err := srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}