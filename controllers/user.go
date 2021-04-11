package controllers

import (
	"fmt"
	"gogal/helpers"
	"gogal/models"
	"gogal/services"
	"gogal/views"
	"log"
	"net/http"
)

func NewUser(us *services.UserService) *User {
	return &User{
		newView:   views.NewView("register"),
		loginView: views.NewView("login"),
		us:        us,
	}
}

type User struct {
	newView   *views.View
	loginView *views.View
	us        *services.UserService
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	u.newView.Render(w, nil)
}

type UserForm struct {
	Name     string `schema:"email"`
	Password string `schema:"password"`
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	user, err := userFromRequest(r)

	if err != nil {
		log.Panic(err)
	}

	err = u.us.Create(user)

	if err != nil {
		fmt.Fprint(w, "Server error")
		log.Println(err.Error())
		return
	}
	fmt.Fprint(w, user)
}

func userFromRequest(r *http.Request) (*models.User, error) {
	var userForm UserForm
	err := helpers.ParseForm(&userForm, r)
	if err != nil {
		log.Panic(err.Error())
	}
	user := &models.User{
		Name:    userForm.Name,
		Pasword: userForm.Password,
	}
	return user, err
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	u.loginView.Render(w, nil)
}

func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	user, err := userFromRequest(r)
	if err != nil {
		log.Panic(err)
	}
	fmt.Fprint(w, user)
}
