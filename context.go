package main

import "net/http"

type Context struct {
	u *User
}

type User struct {
	Name string
}

type handler func(http.ResponseWriter, *http.Request, Context) error

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{}
	err := h(w, r, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
