package main

import (
	"log"
	"os"

	"github.com/zavad4/lab3go/tree/main/server/forums"
)

func ComposeApiServer(port HttpPortNumber) *ApiServer {
	server := &ApiServer{
		Port:   port,
		router: ComposeRouter(),
	}
	return server
}

func ComposeForumsHandler() forums.HttpHandlerFunc {
	db, err := NewDbConnection()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	store := forums.NewStore(db)
	httpHandlerFunc := forums.HttpHandler(store)
	return httpHandlerFunc
}

func ComposeRouter() *Router {
	var routes = map[string]HttpHandlerFunc{
		"/forums": HttpHandlerFunc(ComposeForumsHandler()),
		"/users":  HttpHandlerFunc(ComposeForumsHandler()),
	}
	router := &Router{routes}
	return router
}
