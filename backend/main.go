package main

import(
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"portofolio.com/router"
    "github.com/rs/cors"
)

func main(){
	muxRouter := mux.NewRouter()

	NewRouter := router.AddRouter(muxRouter)

	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    // Use the CORS handler
    handler := c.Handler(NewRouter)

	log.Println("Server serve at port : 8080")
	http.ListenAndServe(":8080", handler)
}