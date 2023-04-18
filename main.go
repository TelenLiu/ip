package main

import (
	"github.com/TelenLiu/ip/api"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	// Setup all routes.  We only service API requests, so this is basic.
	router := httprouter.New()
	router.GET("/", api.GetIP)

	// Setup 404 / 405 handlers.
	router.NotFound = http.HandlerFunc(api.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(api.MethodNotAllowed)

	// Setup middlewares.  For this we're basically adding:
	//	- Support for CORS to make JSONP work.
	handler := cors.Default().Handler(router)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting HTTP server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
