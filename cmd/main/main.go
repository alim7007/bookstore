package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alim7007/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("server listening on port:8080")
	if err := http.ListenAndServe("localhost:8010", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}