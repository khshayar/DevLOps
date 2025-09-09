package main

import (
	"fmt"
	"net/http"
)
const port = ":8585"

func main(){

	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)

	fmt.Println("Starting server at port",port)
	http.ListenAndServe(port,nil)

}
