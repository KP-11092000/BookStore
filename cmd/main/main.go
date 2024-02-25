package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/KP-11092000/Bookstore/pkg/routers"
)
func main(){
	r:= mux.NewRouter()
	routers.RegisterBookStoreRouters(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("local:host:9010",r))
}