package datasvc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kim3z/icat-project-work/pkg/datastreamtypes"

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

// RegisterHandlers register http handlers for bloodpressure
func RegisterHandlers(router *mux.Router, service Service) {
	res := resource{service}

	router.HandleFunc("", res.index).Methods("GET")

	router.HandleFunc("/all/blood-pressure", res.allBloodPressure).Methods("GET")
	router.HandleFunc("/all/temperature", res.allTemperature).Methods("GET")

	router.HandleFunc("/current/temperature", res.currentTemperature).Methods("GET")

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
	res.websocketDataStreamer(&w, r, datastreamtypes.AllBloodPressure)
}

// GET /api/v<x>/data/all/temperature
func (res resource) allTemperature(w http.ResponseWriter, r *http.Request) {
	res.websocketDataStreamer(&w, r, datastreamtypes.AllTemperature)
}

// GET /api/v<x>/data/current/temperature
func (res resource) currentTemperature(w http.ResponseWriter, r *http.Request) {
	res.websocketDataStreamer(&w, r, datastreamtypes.CurrentTemperature)
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

func (res resource) websocketDataStreamer(w *http.ResponseWriter, r *http.Request, dsType datastreamtypes.DataStreamType) {
	ws, err := wsocket.Upgrade(*w, r)
	if err != nil {
		fmt.Fprintf(*w, "%+v\n", err)
	}

	go func(conn *websocket.Conn) {
		for {
			ticker := time.NewTicker(2 * time.Second)

			for range ticker.C {
				results, err := res.getDataStreamData(dsType)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				jsonString, err := json.Marshal(results)

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
					fmt.Println(err.Error())
					return
				}

			}
		}
	}(ws)
}

func (res resource) getDataStreamData(dsType datastreamtypes.DataStreamType) (interface{}, error) {
	switch dsType {
	case datastreamtypes.AllBloodPressure:
		bpData, err := res.service.AllBloodPressure()
		if err != nil {
			return nil, err
		}
		return bpData, nil
	case datastreamtypes.AllTemperature:
		tempData, err := res.service.AllTemperature()
		if err != nil {
			return nil, err
		}
		return tempData, nil
	case datastreamtypes.CurrentTemperature:
		currentTempData, err := res.service.CurrentTemperature()
		fmt.Printf("Current temperature %v\n", currentTempData.Temperature)
		if err != nil {
			return nil, err
		}
		return currentTempData, nil
	default:
		return nil, errors.New("Undefined DataStreamType")
	}
}
