package tools

import(
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct{
	UserName string
	AuthCode string
}
type CoinDetails struct{
UserName string
Coin int
}

type DatabaseInterface interface{
	GetUserLoginDetails(userName string) *LoginDetails
	GetCoinBalance(userName string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (DatabaseInterface,error){
	var database DatabaseInterface=&mockDB{}
	var err error=database.SetupDatabase()
	if err!=nil{
		log.Error(err)
		return nil,err
	}

	return database,nil
}