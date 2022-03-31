package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func SuccessJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

func ErrorJSON(w http.ResponseWriter, statusCode int, err error) {
	SuccessJSON(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}
