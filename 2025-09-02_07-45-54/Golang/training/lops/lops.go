package main

import (
	"fmt"
	
)

func main(){

	//simple loops 

	// for i:=0;i<=10;i++{
	// 	for j:=0;j<=i;j++{
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println(" ")
	// }

	var animals []string =[]string{"dog","cat","humen","lion","dunky"}

	//first i range
	for i,animal:=range animals{
		fmt.Println(i,animal)
	}
	//secound 
	for _,animal:=range animals{
		fmt.Println(animal)
	}

}