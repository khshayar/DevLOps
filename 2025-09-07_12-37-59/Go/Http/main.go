package main

import (
	"errors"
	"fmt"
	"net/http"

)

const port = ":8585"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about",About)
	http.HandleFunc("/dividing",Divide)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "fuck  World!")
}

func About(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w ,"this is my about body")
}

func Divide(w http.ResponseWriter , r *http.Request){
	v ,err := divideValue(100.0 , 0.0 )
	if(err != nil){
		fmt.Fprintf(w,"%s",err)
	}else{
		fmt.Fprintf(w,"%f",v)
	}
	
}

func divideValue(x,y float32) (float32,error){
	
	if(y <= 0){
		err:= errors.New("we cannot dividing")
		return 0,err
	}else{
		return x/y ,nil
	}
}