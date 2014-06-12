package main

import (
	"net/http"

	"appengine/user"

	"appengine"
)

type Context struct {
	c appengine.Context
	u *user.User
}

type handler func(http.ResponseWriter, *http.Request, Context) error

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{}
	ctx.c = appengine.NewContext(r)
	ctx.u = user.Current(ctx.c)
	err := h(w, r, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
