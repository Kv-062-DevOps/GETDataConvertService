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
	metrics.Count()
	metrics.Hist()
	//metrics.Collect()
	router := routes.NewRouter()
	//metrics.Output()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HOST_PORT"), router))

}
