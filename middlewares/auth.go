package middlewares

import (
	"gogal/helpers"
	"gogal/services"
	"net/http"
)

type Auth struct {
	us services.AuthService
}

func NewAuthMiddelware(us *services.AuthService) *Auth {
	return &Auth{
		us: *us,
	}
}

func (am *Auth) IsLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember")
		if err != nil {
			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		user, err := am.us.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		contx := r.Context()
		ctx := helpers.SetUserToContext(contx, user)
		r = r.WithContext(ctx)
		next(rw, r)

	})
}
