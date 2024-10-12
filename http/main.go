package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//variable that declares the url
const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
 fmt.Println("Network Requests Demo")
 response, err := http.Get(url)
 if err != nil {
 panic(err)
 } 
 fmt.Printf("Response Type: %T\n", response)
 defer response.Body.Close()
 databytes, err := ioutil.ReadAll(response.Body)
 if err != nil {
 panic(err)
 }
 content := string(databytes)
 fmt.Println(content)
}