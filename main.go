package main

import (
	"GETDataConvertService/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(os.Getenv("HOST_PORT"), router))

}
