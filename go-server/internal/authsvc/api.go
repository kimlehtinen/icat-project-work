package authsvc

import (
	"encoding/json"
	"net/http"

	"github.com/kim3z/icat-project-work/pkg/errorhandler"

	"github.com/gorilla/mux"
)

type resource struct {
	service Service
}

type message struct {
	Message string `json:"message"`
}

// RegisterHandlers register http handlers for auth
func RegisterHandlers(router *mux.Router, service Service) {
	res := resource{service}

	router.HandleFunc("", res.index).Methods("GET")
	router.HandleFunc("/register", res.register).Methods("POST")
	router.HandleFunc("/login", res.login).Methods("POST")
}

// GET /api/auth
func (res resource) index(w http.ResponseWriter, r *http.Request) {
	message := message{
		Message: "API auth service on port 8081",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

// POST /api/auth/register
func (res resource) register(w http.ResponseWriter, r *http.Request) {
	var input AuthUser
	_ = json.NewDecoder(r.Body).Decode(&input)
	_, err := res.service.Register(input)
	if err != nil {
		errorhandler.NewJsonErrorMessage(w, http.StatusBadRequest, err)
		return
	}

	jwtToken, err := res.service.Auth(input)

	if err != nil {
		errorhandler.NewJsonErrorMessage(w, http.StatusUnauthorized, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(jwtToken); err != nil {
		panic(err)
	}
}

// POST /api/auth/login
func (res resource) login(w http.ResponseWriter, r *http.Request) {
	var input AuthUser
	_ = json.NewDecoder(r.Body).Decode(&input)
	userCreated, err := res.service.Auth(input)
	if err != nil {
		errorhandler.NewJsonErrorMessage(w, http.StatusUnauthorized, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userCreated); err != nil {
		panic(err)
	}
}
