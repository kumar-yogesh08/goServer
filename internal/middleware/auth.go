package middleware

import (
	"net/http"
	"errors"
)
err:=errors.New("Invalid User Name or Token")
func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userName string=r.URL.Query().Get("userName")
		var authCode string=r.Header.Get("Autorization")
	
		if userName== "" || authCode== "" {
			
		}
	})

}