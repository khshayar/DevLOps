package main

import (
	"fmt"
	"unsafe"
)

type Orde struct {
	name string
	ordestustus Ordestustus
	id int}

const apiurl="https://api.example.com/data"

type Ordestustus int

const (
	one Ordestustus = iota
	two
	three	)


func main() {

	var num int = 10
	order1:= Orde{id: 1, name: "Order1", ordestustus:three}
	
	refren(num)
	fmt.Printf("The value of num is %d and num byte is %d \n",num,unsafe.Sizeof(num))
	fmt.Println(apiurl)
	fmt.Println(one, two, three)
	fmt.Println(order1)

	i := 10
	ip:= &i
	fmt.Printf("is %d and ip is %d and %d and %d\n", i,&i,*ip,ip)



}	

func refren(x int ){
	x=x + 10
}