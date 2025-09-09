package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8585"

func main(){

	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)

	_= http.ListenAndServe(port,nil)
	
}

func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err:=parsedTemplate.Execute(w,nil)
	if err != nil{
		fmt.Println("Error parsing template:",err)
		return
	}
	
}