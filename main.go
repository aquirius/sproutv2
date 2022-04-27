package main

import (
	"log"
	"os"
	"sprout/m/v2/internal/systems/core"
	"sprout/m/v2/internal/systems/login"

	//"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	httpAddr          = getenv("SPRT_ADDR", "localhost:1337")
	mysqlDSN          = getenv("SPRT_MYSQL", "sprout:sprout@tcp(localhost)/sprout")
	urlPrefixBackend  = getenv("SPRT_URL_PREFIX_BACKEND", "http://localhost:1337")
	urlPrefixFrontend = getenv("SPRT_URL_PREFIX_FRONTEND", "http://localhost:7331")
	fsPath            = getenv("SPRT_FS_PATH", "uploads")
)

func getenv(key string, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func connectSQL(dsn string) *sqlx.DB {
	log.Print("connecting with mysql...")

	return sqlx.MustConnect("mysql", mysqlDSN)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Runtime struct {
	db    sqlx.DB
	core  core.Core
	login login.Login
}

func BuildRuntime() Runtime {
	// create uploads folder if it doesn't exist
	if err := os.MkdirAll("uploads", 0777); err != nil {
		panic(err)
	}
	sql := connectSQL(mysqlDSN)

	coreProvider := core.NewCoreProvider(sql, "sql")
	loginProvider := login.NewLoginProvider()
	return Runtime{
		db:    *sql,
		core:  *coreProvider.Core,
		login: *loginProvider.Login,
	}
}

func main() {
	BuildRuntime()
}
