package web

import "net/http"

type AdminAuth struct {
	User string
	Pass string
}

func NewAdminAuth(user, pass string) *AdminAuth {
	return &AdminAuth{User: user, Pass: pass}

}

func (a *AdminAuth) Require(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != a.User || pass != a.Pass {
			w.Header().Set("WWW-Authenticate", `Basic realm="admin"`)

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
