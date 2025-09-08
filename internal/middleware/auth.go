package middleware

import (
	"net/http"
	"errors"
	log "github.com/sirupsen/logrus"
	"goServer/internal/tools"
	"goServer/api"
)
var UnAuthError=errors.New("Invalid User Name or Token")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userName string=r.URL.Query().Get("userName")
		var authCode string=r.Header.Get("Autorization")
	
		if userName== "" || authCode== "" {
			log.Error(UnAuthError)
			api.RequestErrorHandler(w,UnAuthError)
			return
		}

		
		var database tools.DatabaseInterface
		database,err:=tools.NewDatabase()
		if err!=nil{
			api.IntenalErrorHandler(w)
			return 
		}
		
		loginDetails:=(database).GetUserLoginDetails(userName)

		if (loginDetails==nil || authCode!=(loginDetails).AuthCode){
			log.Error(UnAuthError)
			api.RequestErrorHandler(w,UnAuthError)
			return 
		}
		next.ServeHTTP(w,r)
})

}