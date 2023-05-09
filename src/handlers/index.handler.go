package handlers

import (
	"fmt"
	"net/http"
	"webScraper/src/constants"
)

func Index(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(constants.SESSION_TOKEN)
	w.Write([]byte(fmt.Sprintf("Bienvenido al web scraping usuario %s", userId)))
}
