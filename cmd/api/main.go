package main
import (
"fmt"
"net/http"
"goServer/internal/handler"
log "github.com/sirupsen/logrus"
"github.com/go-chi/chi"

)

func main(){
	log.SetReportCaller(true)

	var r *chi.Mux= chi.NewRouter()
	
	handler.Handler(r)

	fmt.Println("Go Server started at 8080")
	
	err:=http.ListenAndServe(":8080",r)
	
	if err!=nil{
		log.Error(err)
	}
}