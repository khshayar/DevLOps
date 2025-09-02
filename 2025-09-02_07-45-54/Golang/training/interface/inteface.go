package main

import "log"

type Animal interface{
	Say() string
	leg() int
}
func main ( ){

	p1:=humen{
		Name: "ll",Age: 21,
	}
	c1:=cat{Name: "pishi",Age: 1}

	Information(&p1)
	Information(&c1)
}

type humen struct{
	Name string
	Age int
}

type cat struct{
	Name string
	Age int
}

func (p *humen)Say()string{
	return "fuck you"
}

func(c *cat) Say() string{
	return "miaooo"
}

func(p *humen) leg() int{
	return 4
}
func(p *cat) leg() int{
	return 2
}

func Information(a Animal){
	log.Println(a.Say(),a.leg())
}