package main

import "log"

func main(){
	// istrue:=false

	// if istrue{
	// 	log.Println("is true is ",istrue)
	// }else{
	// 	log.Println("is true is ",istrue)
	// }

	day:="s"

	switch day {
	case "sunday":
		log.Println("this is the first of the start day")

	case "monday":
		log.Println("this is the secound  of the day")
	case "tuseday":
		log.Println("this is a third of the day")
	default:
		log.Println("out off the range")
	}



}