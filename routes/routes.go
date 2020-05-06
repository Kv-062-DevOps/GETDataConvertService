package routes

import (
	"GETDataConvertService/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	var route Route
	route.Name = "Root"
	route.Method = "GET"
	route.Pattern = "/"
	route.HandlerFunc = handlers.RootHandler
	router := mux.NewRouter().StrictSlash(true)

	// setting route for metrics output
	var metricsRoute Route
	metricsRoute.Name = "metrics"
	metricsRoute.Method = "GET"
	metricsRoute.Pattern = "/metrics"

	router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)

	// handler for metrics
	router.Methods(metricsRoute.Method).Path(metricsRoute.Pattern).Name(metricsRoute.Name).Handler(promhttp.Handler())
	return router
}
