package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, lets learn web requests")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl="http://localhost:5500/get"
	response,err:=http.Get(myurl)
	if err!=nil{
		panic(err)
	}
	
	defer response.Body.Close()
	fmt.Println("The response is:",response.StatusCode)
}