package middleware

import (
	"golang-beginner-21/utils"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")
		if authHeader != "" {
			utils.RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})

}
