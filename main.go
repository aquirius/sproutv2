package main

import (
	"net/http"
	"os"
	"sprout/m/v2/internal/server"
	"sprout/m/v2/internal/systems/core"
	"sprout/m/v2/internal/systems/plant"
	"sprout/m/v2/internal/systems/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Runtime struct {
	db     *sqlx.DB
	server *server.Server
	core   *core.Core
	plant  *plant.PlantSystem
	user   *user.UserSystem
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

	plantProvider := plant.NewPlantProvider(&server.Sql)
	plant := plantProvider.NewPlant()

	userProvider := user.NewUserProvider(server.Sql)
	user := userProvider.NewUser()

	return Runtime{
		db:     &server.Sql,
		server: server,
		core:   core,
		plant:  plant,
		user:   user,
	}
}
func main() {
	rt := BuildRuntime()
	mux := http.NewServeMux()

	userH := rt.user
	plantH := rt.plant
	mux.Handle("/login", userH)
	mux.Handle("/register", userH)
	mux.Handle("/plants", plantH)

	http.ListenAndServe(":1234", mux)
}
