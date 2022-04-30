package login

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type LoginHandler struct{}

// LoginProvider provides *Login
type LoginProvider struct {
	Login *Login
}

// Login is capable of providing login access
type Login struct {
	db sqlx.DB
}

func (login *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		fmt.Println("hit")
		return
	case r.Method == http.MethodPost:
		return
	default:
		return
	}
}

// NewLoginProvider returns a new Login provider
func NewLoginProvider(db *sqlx.DB) *LoginProvider {
	return &LoginProvider{
		&Login{db: *db},
	}
}

func (b *LoginProvider) NewLogin() *Login {
	return b.Login
}
