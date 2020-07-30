package iotsvc

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/blood-pressure/store", res.storeBloodPressure).Methods("POST")
	router.HandleFunc("/temperature/store", res.storeTemperature).Methods("POST")
}

// GET /api/v<x>/iot
func (res resource) index(w http.ResponseWriter, r *http.Request) {
	message := message{
		Message: "API iot service on port 8082",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

// POST /api/v<x>/iot/blood-pressure/store
func (res resource) storeBloodPressure(w http.ResponseWriter, r *http.Request) {
	diastolic, err := strconv.ParseFloat(r.FormValue("diastolic"), 64)
	if err != nil {
		panic(err)
	}

	systolic, err := strconv.ParseFloat(r.FormValue("systolic"), 64)
	if err != nil {
		panic(err)
	}

	pulsePerMin, err := strconv.ParseFloat(r.FormValue("pulse_per_min"), 64)
	if err != nil {
		panic(err)
	}

	input := CreateBloodPressureRequest{
		Diastolic:   diastolic,
		Systolic:    systolic,
		PulsePerMin: pulsePerMin,
	}

	bpCreated, err := res.service.StoreBloodPressure(input)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(bpCreated); err != nil {
		panic(err)
	}
}

// POST /api/v<x>/iot/temperature/store
func (res resource) storeTemperature(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.FormValue("temperature"), 64)
	if err != nil {
		panic(err)
	}

	temperatureCreated, err := res.service.StoreTemperature(temperature)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(temperatureCreated); err != nil {
		panic(err)
	}
}
