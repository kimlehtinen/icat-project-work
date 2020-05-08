package datasvc

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
	router.HandleFunc("/all", res.all).Methods("GET")
	router.HandleFunc("/find/{id}", res.find).Methods("GET")
	router.HandleFunc("/create", res.create).Methods("POST")
}

// GET /api/data
func (res resource) index(w http.ResponseWriter, r *http.Request) {
	message := message{
		Message: "API data service on port 8080",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

// GET /api/data/all
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

// GET /api/data/find/{id}
func (res resource) find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	bpResult, err := res.service.Find(id)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bpResult); err != nil {
		panic(err)
	}
}

// POST /api/data/create
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
