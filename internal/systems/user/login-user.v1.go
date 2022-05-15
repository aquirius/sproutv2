package user

import (
	"database/sql"
	"fmt"
	"encoding/json"
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

//LoginUser
func (l *User) LoginUserV1(p *LoginUserV1Params, res *LoginUserV1Result) error {
	var accountID *uint64
	if err := l.db.Get(&accountID, `SELECT id FROM user WHERE display_name = ? AND password_hash = ?`, p.Username, p.Password); err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Print("error: ", err)
	}
	if accountID != nil {
		res.Token = "token:set"
	}
	return nil
}

func (l *User) GetLoginUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &LoginUserV1Params{}
	res := &LoginUserV1Result{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var resBody []byte

	err := json.Unmarshal(reqBody, req)
	if err != nil {
		log.Fatalf("user login error")
		return
	}
	err = l.LoginUserV1(req, res)
	fmt.Println(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resBody, _ = json.Marshal("no-user-found")
		return
	}
	resBody, err = json.Marshal(res.Token)
	if err != nil {
		log.Fatalf("marshal")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write(resBody)
}
