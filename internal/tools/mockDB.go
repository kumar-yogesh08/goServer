package tools

import "time"

type mockDB struct{}

var mockLoginData=map[string]LoginDetails{

	"alex":{
		AuthCode:"blahahah",
		UserName: "alex",
	},
	"drup":{
		AuthCode:"maiyage",
		UserName: "drup",
	},
	"mopi":{
		AuthCode:"sadhah",
		UserName:"mopi",
	},
}
var mockCoinData=map[string]CoinDetails{
	"alex":{
		UserName:"alex",
		Coin: 24,
	},
		"lipa":{
		UserName:"lipa",
		Coin: 45,
	},
		"mopi":{
		UserName:"mopi",
		Coin: 78,
	},
}

func (d *mockDB) GetUserLoginDetails(userName string) *LoginDetails{
	time.Sleep(time.Second*1)

	var clientData=LoginDetails{}

	clientData,ok:=	mockLoginData[userName]

	if !ok{
		return nil
	}
	return &clientData

}

func (d *mockDB) GetCoinBalance(userName string) *CoinDetails{
	time.Sleep(time.Second*1)

	var clientData=CoinDetails{}

	clientData,ok:=	mockCoinData[userName]

	if !ok{
		return nil
	}
	return &clientData


}

func (d *mockDB) SetupDatabase() error{
	return nil
}