package main

import "log"

func main(){
	var namelist []string
	namelist=append(namelist, "mary")
	namelist = append(namelist, "ehsan")


	log.Println(namelist)
}
