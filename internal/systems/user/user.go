package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserHandler struct{}

// LoginProvider provides *Login
type UserProvider struct {
	User *User
}

// Login is capable of providing login access
type User struct {
}

func (user *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		fmt.Println("get user")
		user.GetCreateCustomer(w, r)
		return
	case r.Method == http.MethodPost:
		fmt.Println("post user")
		user.GetCreateCustomer(w, r)
		return
	default:
		return
	}
}

func (user *User) GetCreateCustomer(w http.ResponseWriter, r *http.Request) {
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
func NewUserProvider() *UserProvider {

	return &UserProvider{
		&User{},
	}
}

func (b *UserProvider) NewUser() *User {
	return b.User
}
