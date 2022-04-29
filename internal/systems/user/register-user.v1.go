package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// User
type RegisterUser struct {
	Email           string `json:"email"`
	DisplayName     string `json:"display_name"`
	Title           string `json:"title"`
	Salutation      string `json:"salutation"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	DisplayLanguage string `json:"language"`
	Country         string `json:"country"`
	Password        string `json:"password"`
}

type RegisterUserV1Params struct {
	Email           string `json:"email"`
	DisplayName     string `json:"display_name"`
	Title           string `json:"title"`
	Salutation      string `json:"salutation"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	DisplayLanguage string `json:"language"`
	Country         string `json:"country"`
	Password        string `json:"password"`
}

type RegisterUserV1Result struct {
	User *RegisterUser `json:"user"`
}

//CreateCustomerV1 creates a customer object with given arguments
func (l *User) RegisterUserV1(p *RegisterUserV1Params, res *RegisterUserV1Result) error {

	user := &RegisterUser{
		Email:           "asdf@poo.com",
		DisplayName:     "poo",
		Title:           "sir",
		Salutation:      "mr",
		FirstName:       "poo",
		LastName:        "pimpel",
		DisplayLanguage: "ger",
		Country:         "de",
	}
	res = &RegisterUserV1Result{User: user}
	return nil
}

func (l *User) GetRegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	res := &RegisterUserV1Result{}
	req := &RegisterUser{}

	err := l.RegisterUserV1(nil, res)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, req)
	jsonBytes, err := json.Marshal(res)
	fmt.Println(res.User)
	fmt.Println()

	if err != nil {
		log.Fatal("error in json")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
