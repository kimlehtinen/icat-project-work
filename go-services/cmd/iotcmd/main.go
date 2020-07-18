package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/internal/datasvc"
	"github.com/kim3z/icat-project-work/internal/iotsvc"
	"github.com/kim3z/icat-project-work/pkg/dbcontext"
)

var serverPortNumber = 8082

func main() {
	db, err := dbcontext.NewConnection()

	if err != nil {
		log.Println("iotsvc failed to connect to mongodb")
		panic(err)
	}

	router := mux.NewRouter()

	iotsvcRouter := router.PathPrefix("/api/v1/iot").Subrouter()

	iotRepo := iotsvc.InitRepository(db)
	dataRepo := datasvc.InitRepository(db)

	iotService := iotsvc.InitService(iotRepo, dataRepo)

	iotsvc.RegisterHandlers(iotsvcRouter, iotService)

	port := fmt.Sprintf(":%s", strconv.Itoa(serverPortNumber))
	http.ListenAndServe(port, router)
}
