package controllers

import (
	"gogal/views"
	"net/http"
)

func NewUser() *User {
	return &User{
		newView:   views.NewView("views/register.gohtml"),
		loginView: views.NewView("views/login.gohtml"),
	}
}

type User struct {
	newView   *views.View
	loginView *views.View
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	u.newView.Render(w, nil)
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	u.newView.Render(w, nil)
}
