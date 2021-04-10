package controllers

import (
	"fmt"
	"gogal/helpers"
	"gogal/views"
	"log"
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
	u.newView.Render(w, nil)
}

type UserForm struct {
	Name     string `schema:"email"`
	Password string `schema:"password"`
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	var userForm UserForm
	err := helpers.ParseForm(&userForm, r)
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Fprint(w, userForm)
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	u.loginView.Render(w, nil)
}
