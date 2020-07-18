package errorhandler

import (
	"encoding/json"
	"net/http"
)

func NewJsonErrorMessage(w http.ResponseWriter, status int, err error) {
	resp := map[string]interface{}{
		"status":  status,
		"message": err.Error(),
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}
