package main

import (
	"html/template"
	"net/http"
)

var mainPage *template.Template

type Page struct {
	Title   string
	Content interface{}
}

type User struct {
	Name string
}

func init() {
	layout := template.Must(template.ParseFiles("views/layout.html"))
	mainPage = template.Must(layout.Clone())
	mainPage = template.Must(mainPage.ParseFiles("views/main.html"))

	http.HandleFunc("/", rootHandler)
}

func main() {
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		panic(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	mainPage.Execute(w, Page{
		Title:   "Welcome",
		Content: User{Name: "Tom"},
	})
}
