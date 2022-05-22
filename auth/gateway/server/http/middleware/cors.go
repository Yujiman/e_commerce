package middleware

import (
	"net/http"
)

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, X-Satrap-1, X-Satrap-2, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Expose-Headers", "X-Satrap-1, X-Satrap-2")

		if r.Method == "OPTIONS" {
			w.WriteHeader(244)
			return
		}

		next.ServeHTTP(w, r)
	}
}
