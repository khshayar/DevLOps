package main

import (
	"log"

	"github.com/khshayar/GoLearning/helper"
)

func main (){
	log.Println("Hello")

	h1:=helper.Person{
		Name: "eli",
		Age: 12,
		Phone: "357865416",
	}
	log.Println(h1)
}