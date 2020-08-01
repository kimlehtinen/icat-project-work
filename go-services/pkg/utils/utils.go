package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteHttpJson(w http.ResponseWriter, jsonData interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(jsonData); err != nil {
		fmt.Println(err.Error())
	}
}

func WriteHttpJsonError(w http.ResponseWriter, httpStatus int, err error) {
	fmt.Println(err.Error())
	resp := map[string]interface{}{
		"status":  httpStatus,
		"message": err.Error(),
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}
