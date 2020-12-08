package server

import (
	"net/http"
	"time"
)

func (ser *Server) setCookies(w http.ResponseWriter, r *http.Request, account, password string) {
	expire := time.Now()
	expire = expire.AddDate(365, 0, 0)

	cookie := http.Cookie{Name: account, Value: password, Expires: expire}
	http.SetCookie(w, &cookie)
}

func (ser *Server) getCookies(w http.ResponseWriter, r *http.Request) (string, string) {
	if cookie, err := r.Cookie("account"); err == nil {
		return cookie.Name, cookie.Value
	}
	return "none", "none"
}
