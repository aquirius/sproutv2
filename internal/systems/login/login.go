package login

import (
	"fmt"
	"net/http"
)

type LoginHandler struct{}

// LoginProvider provides *Login
type LoginProvider struct {
	Login *Login
}

// Login is capable of providing login access
type Login struct {
}

func (login *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		fmt.Println("hit")
		login.GetCreateCustomerHandler(w, r)
		return
	case r.Method == http.MethodPost:
		login.GetCreateCustomerHandler(w, r)
		return
	default:
		return
	}
}

// NewLoginProvider returns a new Login provider
func NewLoginProvider() *LoginProvider {
	return &LoginProvider{
		&Login{},
	}
}

func (b *LoginProvider) NewLogin() *Login {
	return b.Login
}
