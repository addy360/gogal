package middlewares

import (
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

		_, err = am.us.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next(rw, r)

	})
}
