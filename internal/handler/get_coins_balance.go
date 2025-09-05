package handler

import (
	"encoding/json"
	"goServer/api"
	"goServer/internal/tools"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {

	var params = api.CoinBalanceParam{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(params, r.URL.Query())
	if err != nil {
		log.Error(err.Error())
		api.IntenalErrorHandler(w)
		return
	}
	var database *tool.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err.Error())
		api.IntenalErrorHandler(w)

	}

	var tokenDetails *tools.CoinDetails

	tokenDetails = (*database).GetCoinBalance(params.UserName)
	if tokenDetails == nil {
		log.Error(err)
		api.IntenalErrorHandler(w)
		return
	}
	var response = api.CoinBalanceResponse{
		Balance: int64((*tokenDetails).Coin),
		Code:    http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err.Error())
		api.IntenalErrorHandler(w)
		return
	}

}
