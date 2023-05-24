package handlers

import (
	"fmt"
	"log"
	"net/http"
	"webScraper/src/constants"
)

func Index(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(constants.SESSION_TOKEN)
	_, err := w.Write([]byte(fmt.Sprintf("Bienvenido al web scraping usuario %s", userId)))
	if err != nil {
		log.Println("Error sending the response ", err)
		return
	}
}
