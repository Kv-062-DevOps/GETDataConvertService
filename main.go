package main

import (
	"GETDataConvertService/routes"
	"log"
	"net/http"
	"os"

	"github.com/Kv-062-DevOps/monitoring/metrics"
)

//testing GitHub Action again
func main() {

	// call RegMetric() func from metrics module
	metrics.RegMetrics()

	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HOST_PORT"), router))

}
