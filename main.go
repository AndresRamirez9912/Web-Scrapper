package main

import (
	"log"
	"net/http"
	"webScraper/src/handlers"

	"github.com/go-chi/chi"
)

func main() {
	// Create router
	router := chi.NewRouter()

	// Set handlers
	router.Get("/exito", handlers.GetExitoData)
	router.Get("/amazon", handlers.GetAmazonData)
	router.Get("/jumbo", handlers.GetJumboData)

	// Start server
	log.Println("Starting Server at port: 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("Error creating the server: ", err)
	}
}
