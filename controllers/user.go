package controllers

import (
	"fmt"
	"gogal/views"
	"log"
	"net/http"

	"github.com/gorilla/schema"
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

type UserForm struct {
	Name     string `schema:"email"`
	Password string `schema:"password"`
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Panic(err.Error())
	}

	dec := schema.NewDecoder()
	var userForm UserForm
	dec.Decode(&userForm, r.PostForm)

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, userForm)
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	u.newView.Render(w, nil)
}
