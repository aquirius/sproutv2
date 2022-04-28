package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sprout/m/v2/internal/systems/core"
	"sprout/m/v2/internal/systems/login"
	"sprout/m/v2/internal/systems/user"

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

type CurrentUser struct {
	customer login.Customer
}

func connectRPC(port int) error {
	login := new(login.Login)
	//core := new(core.Core)

	err := rpc.Register(login)
	if err != nil {
		log.Fatal(err)
	}
	/*err = rpc.Register(core)
	if err != nil {
		log.Fatal(err)
	}*/

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", fmt.Sprintf(":%d", port))

	fmt.Println(l.Addr().String())
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)
	return nil
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
	rpcErr := connectRPC(1234)
	if rpcErr != nil {
		log.Fatal(rpcErr)
	}
	coreProvider := core.NewCoreProvider(sql, "sql")
	loginProvider := login.NewLoginProvider()

	return Runtime{
		db:    *sql,
		core:  *coreProvider.Core,
		login: *loginProvider.Login,
	}
}

func getRPC() {
	client, err := rpc.DialHTTP("tcp", ":1234")

	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	args := &login.CreateCustomerV1Params{}
	var result login.CreateCustomerV1Result
	err = client.Call("Login.CreateCustomerV1", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	log.Printf("asdf", result.Customer.Email)
}

func main() {
	BuildRuntime()
	mux := http.NewServeMux()

	userH := &user.UserHandler{}
	mux.Handle("/users", userH)
	mux.Handle("/", userH)

	http.ListenAndServe(":1234", mux)
}
