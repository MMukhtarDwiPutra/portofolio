package main

import(
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"portofolio.com/router"
)

func main(){
	muxRouter := mux.NewRouter()

	NewRouter := router.AddRouter(muxRouter)

	log.Println("Server serve at port : 8080")
	http.ListenAndServe(":8080", NewRouter)
}