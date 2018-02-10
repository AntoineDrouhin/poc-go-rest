package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Help",
		"GET",
		"/help",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todo",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todo/{todoId}",
		TodoShow,
	},
}
