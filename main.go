package main

import (
	"github.com/qt-sc/server/middleware"
	"log"
	"net/http"

	route "github.com/qt-sc/server/route"
)

const PORT string = "8080"

func main() {

	log.Printf("Server started")

	router := route.NewRouter()

	router.Use(middleware.Auth)
	log.Fatal(http.ListenAndServe(":8080", router))

}