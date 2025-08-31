package main

import "fmt"

func main(){
	fmt.Println("fuck you")

	var sayHi string

	sayHi="fuck your mom"

	fmt.Println(sayHi)

	var i int 

	i = 5
	fmt.Println("the number of i  is ",i)

	what := saySomething()
	
	fmt.Println(what)
}

func saySomething() string {
	return "say something "
}