package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParam struct {
	UserName string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

type Error struct {
	Message string
	code    int
}

func writeError(w http.ResponseWriter, message string, code int) {

	resp := Error{
		code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "apllication/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	IntenalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Server Error", http.StatusInternalServerError)
	}
)
