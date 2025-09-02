package main

import (
	"log"
)

type person struct{
	FirstName string 
	LastName string
}

func main(){

	llsd:= make(map[string] string)
	llsd["name"]="ali"
	llsd["fg"]="fumnf"
	log.Println(llsd["name"])
	smp:= make(map[string] person)

	user:= person{
	FirstName: "ali",
	LastName: "jahani",
	}
	
	smp["user"] = user

	log.Println(smp["user"].FirstName,smp["user"].LastName)
}


