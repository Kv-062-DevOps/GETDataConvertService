package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"GETDataConvertService/handlers"
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
	router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	return router
}
