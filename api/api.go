package api

import(
	"net/http"
	"encoding/json"
)
type CoinBalanceParam struct{
	UserName string
}

type CoinBalanceResponse struct{
	code int
	Balance int64
}

type Error struct{

	Message string
	code int

}

func writeError(w http.ResponseWriter,message string,code int){

	resp:=Error{
		code:code,
		Message: message
	}

	w.Header().Set("Content-Type","apllication/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

