package systems

import (
	"net/http"
	"sprout/m/v2/internal/systems/core"
	"sprout/m/v2/internal/systems/login"
	"sprout/m/v2/internal/systems/user"
)

type Systems struct {
	login login.Login
	core  core.Core
	user  user.User
}

func (s *Systems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		//s.user.Get(w, r)
		return
	case r.Method == http.MethodPost:
		//s.login.CreateCustomerV1Handler(w, r)
		return
	default:
		return
	}
}
