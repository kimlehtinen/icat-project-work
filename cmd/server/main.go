package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/internal/bloodpressure"
	"github.com/kim3z/icat-project-work/pkg/dbcontext"
)

var serverPortNumber = 8080

func main() {
	db, err := dbcontext.NewConnection()

	if err != nil {
		log.Println("Database connection failed")
		panic(err)
	}

	router := mux.NewRouter()

	// /api
	apiRouter := router.PathPrefix("/api").Subrouter()

	// /api/blood-pressure
	bpRouter := apiRouter.PathPrefix("/blood-pressure").Subrouter()
	bloodpressure.RegisterHandlers(bpRouter, bloodpressure.InitService(bloodpressure.InitRepository(db)))

	port := fmt.Sprintf(":%s", strconv.Itoa(serverPortNumber))
	http.ListenAndServe(port, router)
}
