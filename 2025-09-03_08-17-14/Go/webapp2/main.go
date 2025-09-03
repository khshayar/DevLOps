package main

import (
	"fmt"
	"net/http"
)
const port string =":8081"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"this is a home page")
}
func About(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"this is a about page")
	sum,_ :=dividing(5,6)
	fmt.Fprintf(w,"this a dividing : %d",sum)
}

func dividing(x,y int) (int, error){
	var sum int
	sum=x+y
	var err error 
	err=nil
	return sum ,err
}

func main(){

	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)

	http.ListenAndServe(port,nil)
}

