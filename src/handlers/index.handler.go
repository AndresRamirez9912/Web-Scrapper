package handlers

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Bienvenido al web scraping"))
	if err != nil {
		log.Println("Error sending the response ", err)
		return
	}
}
