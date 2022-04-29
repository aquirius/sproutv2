package server

import (
	"log"
	"os"

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

func (sql *Server) connectSQL() *sqlx.DB {
	log.Print("connecting with mysql...")
	return sqlx.MustConnect("mysql", mysqlDSN)
}
