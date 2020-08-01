package iotsvc

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/pkg/utils"
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

	utils.WriteHttpJson(w, message, http.StatusOK)
}

// POST /api/v<x>/iot/blood-pressure/store
func (res resource) storeBloodPressure(w http.ResponseWriter, r *http.Request) {
	diastolic, err := strconv.ParseFloat(r.FormValue("diastolic"), 64)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusInternalServerError, err)
		return
	}

	systolic, err := strconv.ParseFloat(r.FormValue("systolic"), 64)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusInternalServerError, err)
		return
	}

	pulsePerMin, err := strconv.ParseFloat(r.FormValue("pulse_per_min"), 64)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusInternalServerError, err)
		return
	}

	input := CreateBloodPressureRequest{
		Diastolic:   diastolic,
		Systolic:    systolic,
		PulsePerMin: pulsePerMin,
	}

	bpCreated, err := res.service.StoreBloodPressure(input)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpJson(w, bpCreated, http.StatusCreated)
}

// POST /api/v<x>/iot/temperature/store
func (res resource) storeTemperature(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.FormValue("temperature"), 64)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusInternalServerError, err)
		return
	}

	temperatureCreated, err := res.service.StoreTemperature(temperature)
	if err != nil {
		utils.WriteHttpJsonError(w, http.StatusInternalServerError, err)
	}

	utils.WriteHttpJson(w, temperatureCreated, http.StatusCreated)
}
