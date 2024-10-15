package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello, lets learn web requests")
	//PerformGetRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(){
	const myurl="http://localhost:8000/get"
	response,err:=http.Get(myurl)
	if err!=nil{
		panic(err)
	}
	
	defer response.Body.Close()
	fmt.Println("The response is:",response.StatusCode)
}
func PerformPostJson(){
	const myurl="http://localhost:8000/post"

	//fake json data
	requestBody:=strings.NewReader(`
	{
	"coursename":"Golang"
	"price":"$20"
	"platform":"youtube"
	
	}
	
	`)
	response,err:=http.Post(myurl,"application/json",requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostJsonRequest(){

}
func PerformPostFormRequest(){
	const myurl="http://localhost:8000/postform"
	data:=url.Values{}
	data.Add("firstname","Ianoh")
	data.Add("lastname","ke")
	data.Add("course","IT")
	
	response,err:=http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content) )
}