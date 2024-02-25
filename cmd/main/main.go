package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	"github.com/KP-11092000/BookStore/pkg/routers"
)
func main(){
	r:= mux.NewRouter()
	routers.RegisterBookStoreRouters(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:5500",r))
}