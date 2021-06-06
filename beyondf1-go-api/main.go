package main

import (
	v1 "beyondf1-v2/golang-api/v1"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// routes
	go v1.V1(router.NewRoute().Subrouter(), "/api/v1")
	/*router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})*/
	// server
	const addr = ":8080"
	log.Printf("Start listening on %s ...", addr)
	err := http.ListenAndServe(addr, router)
	log.Fatal(err)
}
