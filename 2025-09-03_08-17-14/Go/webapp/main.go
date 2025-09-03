package main

import (
	"fmt"
	"net/http"
)

func main(){	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n,err :=fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	http.HandleFunc("/about",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"this is about page")
	})

	http.HandleFunc("/welcome",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"welcome")
	})
	_=http.ListenAndServe(":8081", nil)
}