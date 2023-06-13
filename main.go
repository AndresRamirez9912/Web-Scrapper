package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"webScraper/src/database"
	"webScraper/src/handlers"
	"webScraper/src/middleware"
	"webScraper/src/services"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	// Create the file to store my logs
	logFile, err := os.OpenFile("./../priceTracking.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error creating the log file")
	}
	defer logFile.Close()

	// Set the location of the logs (my file and the default Stdout)
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetOutput(logFile)

	// Load .env Variables
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env File")
	}

	// Check If the DB exist
	exist := database.ChecWebScrapingExist()
	if !exist {
		// Create general Connection
		DB, err := database.CreateConnectionToDatabase("")
		if err != nil {
			log.Println("Error creating the connection to the DB", err)
		}

		// Create the tables
		err = database.CreateTables(DB)
		if err != nil {
			log.Println("Error creating the database")
			return
		}
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
	router.Get("/verify", handlers.VerifyEmail)

	// Handlers with Auth
	auth.Get("/exito", handlers.GetExitoData)
	auth.Get("/amazon", handlers.GetAmazonData)
	auth.Get("/jumbo", handlers.GetJumboData)

	// Create the Cron Job
	cron := cron.New()
	_, err = cron.AddFunc("0 * * * *", services.CheckProducts) // Check The prices every hour (at o'clock)
	if err != nil {
		log.Println("Error creating the Cron Job: ", err)
	}
	cron.Start()

	// Start server
	log.Println("Starting Server at port: 3000")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println("Error creating the server: ", err)
	}
}
