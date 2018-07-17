package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type configuration struct {
	Name string
	Text string
}

func main() {
	file :=  ` 
        { "Name" : "李四" , "Text" : "Dadada" }
    `

	decoder := json.NewDecoder(strings.NewReader(file))
	conf := configuration{}
	err:= decoder.Decode(&conf)

	if err != nil{
		fmt.Println("Error: ", err)
	}

	fmt.Printf(conf.Name)
}
