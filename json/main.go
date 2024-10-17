package main

import (
	"encoding/json"
	"fmt"
)

type epl struct {
	name     string
	trophies int
	players  int
}

func main() {
	fmt.Println("json code")
	encodeJson()
}

func encodeJson(){
	league:=[]epl{
		{"Chelsea",23,56},
		{"Manjesta",12,67},
		{"Liverpool",9,25},
		{"Arsenal",12,45},
	}
fmt.Print(league)
finalJson,err:=json.MarshalIndent(league,"","\t")
if err!=nil{
	panic(err)

}
fmt.Printf("%s\n",finalJson)
}