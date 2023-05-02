package main

import (
	"log"
	"net/http"
	"webScraper/src/database"
	"webScraper/src/handlers"

	"github.com/go-chi/chi"
)

func main() {
	// Check If the DB exist
	exist := database.ChecWebScrapingExist()
	if !exist {
		// Create general Connection
		DB, err := database.CreateConnectionToDatabase("")
		if err != nil {
			log.Println("Error creating the connection to the DB", err)
		}

		// Create the tables
		database.CreateTables(DB)
	}
	log.Println("Successfully connection to the database")

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
