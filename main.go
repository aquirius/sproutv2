package main

import (
	"log"
	"net/http"
	"os"
	"sprout/m/v2/internal/systems/plant"
	"sprout/m/v2/internal/systems/user"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Runtime struct {
	sql   *sqlx.DB
	redis *redis.Client
}

func connectSQL() *sqlx.DB {
	log.Print("connecting with mysql...")
	return sqlx.MustConnect("mysql", "sprout:sprout@tcp(localhost:3311)/sprout")
}

func connectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6969",
		Password: "",
		DB:       0,
	})

}

func BuildRuntime(sql *sqlx.DB, redis *redis.Client) Runtime {
	// create uploads folder if it doesn't exist
	if err := os.MkdirAll("uploads", 0777); err != nil {
		panic(err)
	}

	return Runtime{
		sql:   sql,
		redis: redis,
	}
}
func main() {
	rt := BuildRuntime(
		connectSQL(),
		connectRedis(),
	)
	mux := http.NewServeMux()

	//register all system providers
	plantProvider := plant.NewPlantProvider(rt.sql)
	plant := plantProvider.NewPlant()
	userProvider := user.NewUserProvider(rt.sql)
	user := userProvider.NewUser()

	mux.Handle("/login", user)
	mux.Handle("/register", user)
	mux.Handle("/plants", plant)

	http.ListenAndServe(":1234", mux)
}
