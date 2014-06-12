package main

import (
	"html/template"
	"net/http"
)

var mainPage *template.Template

type Page struct {
	Title   string
	User    *User
	Content interface{}
}

func init() {
	layout := template.Must(template.ParseFiles("views/layout.html"))
	mainPage = template.Must(layout.Clone())
	mainPage = template.Must(mainPage.ParseFiles("views/main.html"))

	http.Handle("/", handler(rootHandler))
	http.Handle("/disconnect", handler(disconnectHandler))
}

func rootHandler(w http.ResponseWriter, r *http.Request, c Context) (err error) {
	if c.u == nil {
		// redirect to login?
	}
	c.u = &User{Name: "Tom"}
	mainPage.Execute(w, Page{
		Title:   "Welcome",
		User:    c.u,
		Content: User{Name: c.u.Name},
	})
	return nil
}

func disconnectHandler(w http.ResponseWriter, r *http.Request, c Context) (err error) {
	return nil
}
