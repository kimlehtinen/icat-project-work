package bloodpressure

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type resource struct {
	service Service
}

// RegisterHandlers register http handlers for bloodpressure
func RegisterHandlers(router *mux.Router, service Service) {
	res := resource{service}
	router.HandleFunc("/results", res.all).Methods("GET")
}

func (res resource) all(w http.ResponseWriter, r *http.Request) {
	results, err := res.service.All()

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}
