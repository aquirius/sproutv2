package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserHandler struct{}

// LoginProvider provides *Login
type UserProvider struct {
	User *User
}

// Login is capable of providing login access
type User struct {
	db sqlx.DB
}

func (user *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (user *User) GetCreateUser(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal("")
	if err != nil {
		log.Fatal("error in json")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// NewLoginProvider returns a new Login provider
func NewUserProvider(db sqlx.DB) *UserProvider {

	return &UserProvider{
		&User{db: db},
	}
}

func (b *UserProvider) NewUser() *User {
	return b.User
}
