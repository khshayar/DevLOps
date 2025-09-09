package main

import (
	"fmt"
	"html/template"
	"net/http"
)



func RenderTemplate(w http.ResponseWriter, name string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + name)
	err:=parsedTemplate.Execute(w,nil)
	if err != nil{
		fmt.Println("Error parsing template:",err)
		return
	}
}