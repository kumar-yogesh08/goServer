package api

// import (
// 	"encoding/json"
// 	"net/http"
// )
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

