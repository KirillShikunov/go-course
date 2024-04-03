package auth

import (
	"net/http"
)

const CookieName = "user_token"

func CheckLoggedIn(r *http.Request) (int, bool) {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return 0, false
	}

	userId, valid := DecodeToken(cookie.Value)
	if !valid {
		return 0, false
	}

	return userId, true
}

func SetLoggedIn(w http.ResponseWriter, newToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    newToken,
		Path:     "/",
		HttpOnly: true,
	})
}
