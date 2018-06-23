package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router:=httprouter.New()
	router.Get("/videos/:vid-id", streamHandler)
	router.POSt("/upload/:vid-id", uploadHandler)
	return router
}

func main() {
	r:=RegisterHandlers()
	http.ListenAndServer(":9000", r)
}