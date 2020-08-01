package authsvc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kim3z/icat-project-work/pkg/middleware"
	"github.com/kim3z/icat-project-work/pkg/models"
	"github.com/kim3z/icat-project-work/pkg/utils"

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
	router.Handle("/user", middleware.Auth(http.HandlerFunc(res.user))).Methods("GET")
}

// GET /api/v<x>/auth
func (res resource) index(w http.ResponseWriter, r *http.Request) {
	message := message{
		Message: "API auth service on port 8081",
	}

	utils.WriteHttpJson(w, message, http.StatusOK)
}

// GET /api/auth/user
func (res resource) user(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(models.JwtKey).(models.JwtClaims)

	if !ok {
		fmt.Println("Error")
	}

	user, err := res.service.Find(claims.UserID)

	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusBadRequest, err)
		return
	}

	// don't return password hash
	user.Password = ""

	utils.WriteHttpJson(w, user, http.StatusOK)
}

// POST /api/v<x>/auth/register
func (res resource) register(w http.ResponseWriter, r *http.Request) {
	var input AuthUser
	_ = json.NewDecoder(r.Body).Decode(&input)
	_, err := res.service.Register(input)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusBadRequest, err)
		return
	}

	jwtToken, err := res.service.Auth(input)

	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusUnauthorized, err)
		return
	}

	utils.WriteHttpJson(w, jwtToken, http.StatusCreated)
}

// POST /api/v<x>/auth/login
func (res resource) login(w http.ResponseWriter, r *http.Request) {
	var input AuthUser
	_ = json.NewDecoder(r.Body).Decode(&input)
	userLogin, err := res.service.Auth(input)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusUnauthorized, err)
		return
	}

	utils.WriteHttpJson(w, userLogin, http.StatusOK)
}
