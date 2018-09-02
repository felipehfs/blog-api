package main

import (
	"log"
	"net/http"

	"github.com/felipehfs/blog/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	rt := controller.Adapt(router, controller.WithDB, controller.Logging, controller.Cors)
	router.Handle("/posts", controller.ReadPost()).Methods("GET")
	router.Handle("/posts", controller.CreatePost()).Methods("POST")
	router.Handle("/posts/{id}", controller.UpdatePost()).Methods("PUT")
	router.Handle("/posts/{id}", controller.FindByIDPost()).Methods("GET")
	router.Handle("/posts/{id}", controller.RemovePost()).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8082", rt))
}
