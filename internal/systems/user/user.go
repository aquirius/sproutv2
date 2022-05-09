package user

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserHandler struct{}

// LoginProvider provides *Login
type UserProvider struct {
	UserSystem *UserSystem
}

// Login is capable of providing login access
type UserSystem struct {
	db *sqlx.DB
}

func (user *UserSystem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:

		return
	case r.Method == http.MethodPost:
		if r.URL.Path == "/login" {
			user.GetLoginUserHandler(w, r)
		}
		if r.URL.Path == "/register" {
			user.GetRegisterUserHandler(w, r)
		}
		return
	default:
		return
	}
}

// NewLoginProvider returns a new Login provider
func NewUserProvider(db *sqlx.DB) *UserProvider {

	return &UserProvider{
		&UserSystem{db: db},
	}
}

func (b *UserProvider) NewUser() *UserSystem {
	return b.UserSystem
}
