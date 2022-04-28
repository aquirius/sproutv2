package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginUserV1Params struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginUserV1Result struct {
	Token string
}

//CreateCustomerV1 creates a customer object with given arguments
func (l *Login) LoginUserV1(p *LoginUserV1Params, res *LoginUserV1Result) (error, *string) {
	var customerID *int
	dbRes, err := l.db.Query("select id from customers where display_name = ? and password_hash = ?", p.Username, p.Password)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	var token string
	defer dbRes.Close()
	for dbRes.Next() {
		err := dbRes.Scan(&customerID)
		if err != nil {
			return err, nil
		}
		if customerID != nil {
			token = "token:set"
		}
		fmt.Println(customerID)
	}
	return nil, &token
}

func (l *Login) GetLoginUserHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	req := &LoginUserV1Params{}
	res := &LoginUserV1Result{}

	json.Unmarshal(reqBody, req)
	err, token := l.LoginUserV1(req, nil)
	if err != nil {
		log.Fatalf("user login error")
		return
	}
	res.Token = *token
	jsonBytes, err := json.Marshal(res.Token)
	fmt.Println(jsonBytes)
	if err != nil {
		log.Fatalf("marshal")
		return
	}

	if err != nil {
		log.Fatal("error in json")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
