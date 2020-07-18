package iotsvc

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type resource struct {
	service Service
}

type message struct {
	Message string `json:"message"`
}

// RegisterHandlers register http handlers for bloodpressure
func RegisterHandlers(router *mux.Router, service Service) {
	res := resource{service}

	router.HandleFunc("", res.index).Methods("GET")
	router.HandleFunc("/create", res.create).Methods("POST")
}

// GET /api/v<x>/iot
func (res resource) index(w http.ResponseWriter, r *http.Request) {
	message := message{
		Message: "API iot service on port 8082",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

// POST /api/v<x>/iot/create
func (res resource) create(w http.ResponseWriter, r *http.Request) {
	var input CreateBloodPressureRequest
	_ = json.NewDecoder(r.Body).Decode(&input)
	bpCreated, err := res.service.Create(input)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(bpCreated); err != nil {
		panic(err)
	}
}
