package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/internal/bloodpressure"
	"github.com/kim3z/icat-project-work/pkg/dbcontext"
)

var serverPort = 8080

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

	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", router)
}
