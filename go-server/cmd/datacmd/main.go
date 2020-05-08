package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/internal/datasvc"
	"github.com/kim3z/icat-project-work/pkg/dbcontext"
)

var serverPortNumber = 8080

func main() {
	db, err := dbcontext.NewConnection()

	if err != nil {
		log.Println("datasvc failed to connect to mongodb")
		panic(err)
	}

	router := mux.NewRouter()

	// /api/datasvc
	datasvcRouter := router.PathPrefix("/api/data").Subrouter()
	datasvc.RegisterHandlers(datasvcRouter, datasvc.InitService(datasvc.InitRepository(db)))

	port := fmt.Sprintf(":%s", strconv.Itoa(serverPortNumber))
	http.ListenAndServe(port, router)
}
