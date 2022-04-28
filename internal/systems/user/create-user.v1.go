package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// User
type CreateUser struct {
	CustomerID          uint64 `json:"customer_id"`
	Email               string `json:"email"`
	RegisteredTS        uint64 `json:"registered_ts"`
	LastLoginTS         uint64 `json:"last_login_ts"`
	EmailToken          string `json:"email_toke"`
	EmailConfirmationTS uint64 `json:"email_confirmation_ts"`
	Status              bool   `json:"status"`
	DisplayName         string `json:"display_name"`
	ImageID             string `json:"image_id"`
	CoverID             string `json:"cover_id"`
	Title               string `json:"title"`
	Salutation          string `json:"salutation"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Birthday            string `json:"birthday"`
	DisplayLanguage     string `json:"language"`
	Country             string `json:"country"`
}

type CreateUserV1Params struct {
	CustomerID          uint64 `json:"customer_id"`
	Email               string `json:"email"`
	RegisteredTS        uint64 `json:"registered_ts"`
	LastLoginTS         uint64 `json:"last_login_ts"`
	EmailToken          string `json:"email_toke"`
	EmailConfirmationTS uint64 `json:"email_confirmation_ts"`
	Status              bool   `json:"status"`
	DisplayName         string `json:"display_name"`
	ImageID             string `json:"image_id"`
	CoverID             string `json:"cover_id"`
	Title               string `json:"title"`
	Salutation          string `json:"salutation"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Birthday            string `json:"birthday"`
	DisplayLanguage     string `json:"language"`
	Country             string `json:"country"`
}

type CreateUserV1Result struct {
	User CreateUser `json:"user"`
}

//CreateCustomerV1 creates a customer object with given arguments
func (l *User) CreateUserV1(p *CreateUserV1Params, res *CreateUserV1Result) error {

	user := CreateUser{
		CustomerID:          1,
		Email:               "asdf@poo.com",
		RegisteredTS:        1,
		LastLoginTS:         1,
		EmailToken:          "asdf",
		EmailConfirmationTS: 1,
		Status:              true,
		DisplayName:         "poo",
		ImageID:             "uploads/poo.jpg",
		CoverID:             "uploads/cover/poo.jpg",
		Title:               "sir",
		Salutation:          "mr",
		FirstName:           "poo",
		LastName:            "pimpel",
		Birthday:            "11.1.1112",
		DisplayLanguage:     "ger",
		Country:             "de",
	}
	res = &CreateUserV1Result{User: user}
	return nil
}

func (l *User) GetCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	res := &CreateUserV1Result{}
	err := l.CreateUserV1(nil, res)
	reqBody, _ := ioutil.ReadAll(r.Body)
	req := &CreateUser{}
	json.Unmarshal(reqBody, req)
	jsonBytes, err := json.Marshal(res)
	fmt.Println(res.User)
	fmt.Println(req.Birthday)

	if err != nil {
		log.Fatal("error in json")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
