package main

import "github.com/khshayar/golearn/helper"

const pool = 10

func main(){
	intchan := make(chan int)
	defer close(intchan)

	go sendData(intchan)

	data := <- intchan
	println(data)
}

func sendData(intchan chan int){
	random := helper.Rand(pool)
	intchan <- random
}