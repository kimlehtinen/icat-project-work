package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/internal/bloodpressure"
)

var serverPort = 8080

func main() {
	router := mux.NewRouter()

	bloodpressure.RegisterHandlers(router)

	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", router)
}
