package middleware

import (
	"log"
	"net/http"

	"github.com/Yujiman/e_commerce/basket/gatway/internal/config"
)

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		origin := validateOrigin(r.Header.Get("Origin"))
		log.Println(origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, AccessToken, RefreshToken , Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Expose-Headers", "AccessToken, RefreshToken")

		if r.Method == "OPTIONS" {
			w.WriteHeader(244)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func validateOrigin(origin string) string {
	allowedOrigins := config.GetAllowedCORSOrigin()
	if len(allowedOrigins) == 0 {
		return "*"
	}
	if allowedOrigins[0] == "*" {
		return "*"
	}
	for _, allowed := range allowedOrigins {
		if allowed == origin {
			return origin
		}
	}
	return ""
}
