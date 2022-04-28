package main

import (
	"net/http"
	"os"
	"sprout/m/v2/internal/systems/core"
	"sprout/m/v2/internal/systems/login"
	"sprout/m/v2/internal/systems/user"
	"sprout/m/v2/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Runtime struct {
	db     *sqlx.DB
	server *server.Server
	core   *core.Core
	login  *login.Login
	user   *user.User
}

func BuildRuntime() Runtime {
	// create uploads folder if it doesn't exist
	if err := os.MkdirAll("uploads", 0777); err != nil {
		panic(err)
	}
	serverProvider := server.NewServerProvider()
	server := serverProvider.NewServer()

	coreProvider := core.NewCoreProvider(&server.Sql, "sql")
	core := coreProvider.NewCore()
	loginProvider := login.NewLoginProvider()
	login := loginProvider.NewLogin()
	userProvider := user.NewUserProvider()
	user := userProvider.NewUser()

	return Runtime{
		db:     &server.Sql,
		server: server,
		core:   core,
		login:  login,
		user:   user,
	}
}
func main() {
	rt := BuildRuntime()
	mux := http.NewServeMux()

	userH := rt.user
	loginH := rt.login
	mux.Handle("/users", userH)
	mux.Handle("/login", loginH)
	http.ListenAndServe(":1234", mux)
}
