package main

import (
	"html/template"
	"net/http"

	"appengine/user"
)

var mainPage *template.Template

type Page struct {
	Title   string
	User    *user.User
	Content interface{}
}

type User struct {
	Name string
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
		url, err := user.LoginURL(c.c, r.URL.String())
		if err != nil {
			return err
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return nil
	}
	mainPage.Execute(w, Page{
		Title:   "Welcome",
		User:    c.u,
		Content: User{Name: c.u.Email},
	})
	return nil
}

func disconnectHandler(w http.ResponseWriter, r *http.Request, c Context) (err error) {
	logoutUrl, err := user.LogoutURL(c.c, "/")
	if err != nil {
		return err
	}
	w.Header().Set("Location", logoutUrl)
	w.WriteHeader(http.StatusFound)
	return nil
}
