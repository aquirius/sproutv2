package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// User
type RegisterUser struct {
	Email               string  `json:"email"`
	DisplayName         string  `json:"display_name"`
	Password            string  `json:"password"`
	RegisteredTS        uint64  `json:"registered_ts"`
	LastLoginTS         uint64  `json:"last_login_ts"`
	EmailToken          *string `json:"email_token"`
	EmailConfirmationTS *uint64 `json:"email_confirmation_ts"`
	Status              *bool   `json:"status"`
	ImageID             *string `json:"image_id"`
	CoverID             *string `json:"cover_id"`
	Title               string  `json:"title"`
	Salutation          *string `json:"salutation"`
	FirstName           string  `json:"first_name"`
	LastName            string  `json:"last_name"`
	Birthday            *string `json:"birthday"`
	DisplayLanguage     string  `json:"language"`
	Country             string  `json:"country"`
}

type RegisterUserV1Params struct {
	Email           string  `json:"email"`
	DisplayName     string  `json:"display_name"`
	Password        string  `json:"password"`
	Status          *bool   `json:"status"`
	ImageID         *string `json:"image_id"`
	CoverID         *string `json:"cover_id"`
	Title           string  `json:"title"`
	Salutation      string  `json:"salutation"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Birthday        *string `json:"birthday"`
	DisplayLanguage string  `json:"language"`
	Country         string  `json:"country"`
}

type RegisterUserV1Result struct {
}

//CreateCustomerV1 creates a customer object with given arguments
func (l *UserSystem) RegisterUserV1(p *RegisterUserV1Params, res *RegisterUserV1Result) error {
	fmt.Println(p)
	if p == nil {
		return fmt.Errorf("no params")
	}
	_, err := l.db.Exec(`insert into user (email, display_name, registered_ts, title, salutation, first_name, last_name, language, country, password_hash) values (?,?,?,?,?,?,?,?,?,?)`, p.Email, p.DisplayName, time.Now().Unix(), p.Title, "mr", p.FirstName, p.LastName, p.DisplayLanguage, p.Country, p.Password)
	if err != nil {
		log.Print("error: ", err)
		return err
	}
	return nil
}

func (l *UserSystem) GetRegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	fmt.Println("register")
	req := &RegisterUserV1Params{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, req)
	fmt.Println(req)
	if err != nil {
		log.Fatal("error in json")
		return
	}
	err = l.RegisterUserV1(req, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonBytes, err := json.Marshal("error")
		if err != nil {
			log.Fatal("error in json")
			return
		}
		w.Write(jsonBytes)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
