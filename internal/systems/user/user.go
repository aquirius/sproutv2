package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sprout/m/v2/internal/systems/login"
)

var (
	validIdentifierName = regexp.MustCompile(`^[A-Za-z0-9-]+$`)
	matchMethodName     = regexp.MustCompile(`^(.+)V([0-9]+)$`)
)

type UserHandler struct{}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		h.GetCustomer(w, r)
		return
	case r.Method == http.MethodPost:
		h.GetCustomer(w, r)
		return
	default:
		return
	}
}

func (h *UserHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getcustomer")
	p := &login.CreateCustomerV1Params{}
	res := &login.CreateCustomerV1Result{}
	login.NewLoginProvider().Login.CreateCustomerV1(p, res)

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {

	jsonBytes, err := json.Marshal("")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal("")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

/*func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		internalServerError(w, r)
		return
	}

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}*/

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
