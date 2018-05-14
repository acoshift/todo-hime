package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/acoshift/configfile"
	"github.com/acoshift/hime"
	redisstore "github.com/acoshift/session/store/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"

	"github.com/acoshift/todo-hime/app"
)

func main() {
	config := configfile.NewReader("config")

	db, err := sql.Open("postgres", config.String("sql_db"))
	if err != nil {
		log.Fatal(err)
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal(err)
	}

	sessionHost := config.String("session_host")
	sessionStore := redisstore.New(redisstore.Config{
		Prefix: config.String("session_prefix"),
		Pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", sessionHost)
			},
		},
	})

	himeApp := hime.New()

	err = himeApp.
		TemplateDir("template").
		TemplateRoot("root").
		Minify().
		Handler(app.New(himeApp, app.Config{
			DB:           db,
			Location:     loc,
			SessionStore: sessionStore,
		})).
		GracefulShutdown().
		ListenAndServe(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
