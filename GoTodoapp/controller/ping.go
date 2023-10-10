package controller

import (
	"encoding/json"
	"myapi/views"
	"net/http"
)

// here HandlerFunc used while returning a function
func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}

	}
}
