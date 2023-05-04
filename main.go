package main

import (
	"log"
	"net/http"
	"webScraper/src/database"
	"webScraper/src/handlers"
	"webScraper/src/middleware"

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

	// Create routers
	router := chi.NewRouter()

	// Create a group with the session middleware
	auth := router.Group(nil)
	auth.Use(middleware.CheckAuth)

	// Handlers without Auth
	router.Get("/", handlers.Index)
	router.Post("/register", handlers.Register)
	router.Post("/login", handlers.Login)

	// Handlers with Auth
	auth.Get("/exito", handlers.GetExitoData)
	auth.Get("/amazon", handlers.GetAmazonData)
	auth.Get("/jumbo", handlers.GetJumboData)

	// Start server
	log.Println("Starting Server at port: 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("Error creating the server: ", err)
	}
}
