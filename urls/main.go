package main

import (
	"fmt"
	"net/url"
)


const myurl string="https://github.com/823Njoroge"
func main() {
	fmt.Println("I am learning about urls in golang")
	fmt.Println(myurl)
	result, err:=url.Parse(myurl)
	if err!=nil{
		panic(err)
	}
	fmt.Print("Scheme is:",result.Scheme)
	fmt.Print("Host is:",result.Host)
	fmt.Print("Path is :",result.Path)
	fmt.Print(result.Port())
	fmt.Print(result.RawQuery)
}