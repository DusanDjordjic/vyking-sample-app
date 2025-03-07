package main

import (
	"app/pkg/config"
	"app/pkg/db"
	"app/pkg/router"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port uint
	flag.UintVar(&port, "port", 8080, "specify port to listen on")
	flag.Parse()

	appConfig := config.Parse()
	db.Connect(appConfig)
	log.Printf("connected to database")

	mux := http.NewServeMux()
	router.SetupRouter(mux)

	log.Printf("running server on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
