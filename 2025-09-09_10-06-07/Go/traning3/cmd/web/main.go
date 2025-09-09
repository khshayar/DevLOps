package main

import (
	"fmt"
	"net/http"

	"github.com/khshayr/golearn/pkg/handler"
)
const port = ":8585"

func main(){

	http.HandleFunc("/",handler.Home)
	http.HandleFunc("/about",handler.About)

	fmt.Println("Starting server at port",port)
	http.ListenAndServe(port,nil)

}
