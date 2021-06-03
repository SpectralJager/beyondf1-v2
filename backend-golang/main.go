package main

import (
	"beyondf1-v2/backend-golang/apiv1"
	"beyondf1-v2/backend-golang/data"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	err := data.CreateTables()
	if err != nil {
		os.Exit(1)
	}
	log.Println("Tablse is Created!")
	// router
	router := mux.NewRouter()
	// routes
	// api v1
	go apiv1.Admin_v1(router.NewRoute().Subrouter())
	go apiv1.Site_v1(router.NewRoute().Subrouter())
	//server
	const addr = "0.0.0.0:8000"
	log.Printf("Start listening on %s ...", addr)
	err = http.ListenAndServe(addr, router)
	log.Fatal(err)
}
