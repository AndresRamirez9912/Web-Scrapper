package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"webScraper/src/constants"
	"webScraper/src/database"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entro al Middleware")

		// Check if the session cookie exists
		sessionCookie, err := r.Cookie(constants.SESSION_TOKEN)
		if err != nil {
			log.Println("Client doesn't have session cookie")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		// Check if the session expired
		if !sessionCookie.Expires.Before(time.Now()) {
			log.Println("The cookie expired")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		// Check the session and discover the user
		userId, err := database.GetUserbyCookie(sessionCookie.Value)
		if err != nil {
			log.Println("The user cannot be found with the session key")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		// Send the userId in the context to the handler
		ctx := context.WithValue(r.Context(), constants.USER_ID, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
