package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/internal/authsvc"
	"github.com/kim3z/icat-project-work/pkg/dbcontext"
)

var serverPortNumber = 8080

func main() {
	db, err := dbcontext.NewConnection()

	if err != nil {
		log.Println("authsvc failed to connect to mongodb")
		panic(err)
	}

	router := mux.NewRouter()

	// /api/auth
	authsvcRouter := router.PathPrefix("/api/auth").Subrouter()
	authsvc.RegisterHandlers(authsvcRouter, authsvc.InitService(authsvc.InitRepository(db)))

	port := fmt.Sprintf(":%s", strconv.Itoa(serverPortNumber))
	http.ListenAndServe(port, router)
}
