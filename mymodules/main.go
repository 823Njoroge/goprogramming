package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Hello ")
	greeter()
	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000",r))
}
func greeter(){
	fmt.Println("Mod in golang")
}

func serveHome(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("<h2>Hello and welcome</h2>"))
}