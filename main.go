package main

import (
	"GETDataConvertService/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":30000", router))

}
