package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

func RegisterHandlers() *httprouter.Router {
	log.Printf("preparing to post request\n")
	router:=httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r:=RegisterHandlers()
	http.ListenAndServe(":8000", r)
}