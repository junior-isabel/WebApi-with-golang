package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s %v", r.Method, r.RequestURI, r.Host, time.Now().UTC())
		next(w, r)
	}
}
