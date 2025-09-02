package main

import "fmt"

func main(){

	p:=person{
		Name: "ehsan",
	}
	fmt.Println(p.sayname())
}

type person struct{
	Name string
}

func(p *person) sayname()string{
return p.Name
}