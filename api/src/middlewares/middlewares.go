package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n [%s] %s \n", r.Method, r.RequestURI)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.ErrorJSON(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
