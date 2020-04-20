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
	var route1 Route
	route1.Name = "metrics"
	route1.Method = "GET"
	route1.Pattern = "/metrics"
	router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	router.Methods(route1.Method).Path(route1.Pattern).Name(route1.Name).Handler(promhttp.Handler())
	return router
}
