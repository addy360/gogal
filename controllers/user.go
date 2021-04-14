package controllers

import (
	"fmt"
	"gogal/helpers"
	"gogal/models"
	"gogal/services"
	"gogal/views"
	"net/http"
)

func NewUser(us services.AuthService) *User {
	return &User{
		newView:   views.NewView("register"),
		loginView: views.NewView("login"),
		us:        us,
	}
}

type User struct {
	newView   *views.View
	loginView *views.View
	us        services.AuthService
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	a := services.NewAlert()
	a.Level = a.AlertDanger
	a.Message = "Testing alert levels"
	data := map[string]interface{}{
		"Alert": a,
	}
	u.newView.Render(w, data)
}

type UserForm struct {
	Name     string `schema:"email"`
	Password string `schema:"password"`
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	user, err := userFromRequest(r)

	if err != nil {
		a := services.NewAlert()
		a.Message = err.Error()
		a.Level = a.AlertDanger
		data := map[string]interface{}{
			"Alert": a,
		}
		u.newView.Render(w, data)
	}

	err = u.us.Create(user)

	if err != nil {
		a := services.NewAlert()
		a.Message = err.Error()
		a.Level = a.AlertDanger
		data := map[string]interface{}{
			"Alert": a,
		}
		u.newView.Render(w, data)
	}

	u.us.SignUserIn(user, w)

	http.Redirect(w, r, "/cookie", http.StatusPermanentRedirect)
}

func userFromRequest(r *http.Request) (*models.User, error) {
	var userForm UserForm
	err := helpers.ParseForm(&userForm, r)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Email:   userForm.Name,
		Pasword: userForm.Password,
	}
	return user, nil
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {

	u.loginView.Render(w, nil)
}

func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	user, err := userFromRequest(r)
	if err != nil {
		a := services.NewAlert()
		a.Message = err.Error()
		a.Level = a.AlertDanger
		data := map[string]interface{}{
			"Alert": a,
		}
		u.loginView.Render(w, data)
	}
	_, err = u.us.Authenticate(w, user)
	if err != nil {
		a := services.NewAlert()
		a.Message = err.Error()
		a.Level = a.AlertDanger
		data := map[string]interface{}{
			"Alert": a,
		}
		u.loginView.Render(w, data)
	}

	http.Redirect(w, r, "/cookie", http.StatusPermanentRedirect)
}

func (u *User) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember")
	if err != nil {
		a := services.NewAlert()
		a.Message = err.Error()
		a.Level = a.AlertDanger
		data := map[string]interface{}{
			"Alert": a,
		}
		u.loginView.Render(w, data)
	}

	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		a := services.NewAlert()
		a.Message = err.Error()
		a.Level = a.AlertDanger
		data := map[string]interface{}{
			"Alert": a,
		}
		u.loginView.Render(w, data)
	}

	fmt.Fprint(w, user)
	fmt.Fprint(w, cookie)
}
