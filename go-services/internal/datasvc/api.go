package datasvc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/kim3z/icat-project-work/pkg/middleware"
	"github.com/kim3z/icat-project-work/pkg/wsocket"
)

type resource struct {
	service Service
}

type message struct {
	Message string `json:"message"`
}

func (res resource) writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(2 * time.Second)

		for t := range ticker.C {
			fmt.Printf("In ticker %+v\n", t)

			results, err := res.service.AllBloodPressure()
			if err != nil {
				panic(err)
			}

			jsonString, err := json.Marshal(results)

			if err != nil {
				fmt.Println(err)
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}

// RegisterHandlers register http handlers for bloodpressure
func RegisterHandlers(router *mux.Router, service Service) {
	res := resource{service}

	router.HandleFunc("", res.index).Methods("GET")
	router.HandleFunc("/all/blood-pressure", res.allBloodPressure).Methods("GET")
	router.Handle("/find/{id}", middleware.Auth(http.HandlerFunc(res.find))).Methods("GET")
}

// GET /api/v<x>/data
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

// GET /api/v<x>/data/all/blood-pressure
func (res resource) allBloodPressure(w http.ResponseWriter, r *http.Request) {
	ws, err := wsocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	//go res.writer(ws)
	go func(conn *websocket.Conn) {
		for {
			ticker := time.NewTicker(2 * time.Second)

			for t := range ticker.C {
				fmt.Printf("In ticker %+v\n", t)

				results, err := res.service.AllBloodPressure()
				if err != nil {
					panic(err)
				}

				jsonString, err := json.Marshal(results)

				if err != nil {
					fmt.Println(err)
				}

				if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
					fmt.Println(err)
					return
				}

			}
		}
	}(ws)
}

// GET /api/v<x>/data/find/{id}
func (res resource) find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	bpResult, err := res.service.FindBloodPressure(id)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bpResult); err != nil {
		panic(err)
	}
}
