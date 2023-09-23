package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/acoshift/configfile"
	"github.com/acoshift/pgsql/pgctx"
	_ "github.com/lib/pq"
	"github.com/moonrhythm/hime"
	"github.com/moonrhythm/httpmux"
	"github.com/moonrhythm/parapet"
	"github.com/moonrhythm/session"
	"github.com/moonrhythm/session/store"

	"github.com/acoshift/todo-hime/todoapp"
)

func main() {
	configfile.LoadDotEnv()
	config := configfile.NewEnvReader()

	db, err := sql.Open("postgres", config.String("sql_db"))
	if err != nil {
		log.Fatal(err)
	}

	app := hime.New()
	app.SetServer(parapet.NewBackend()) // optimize config to run behind reverse proxy
	app.ETag = true

	app.Server().UseFunc(session.Middleware(session.Config{
		Rolling:  true,
		Proxy:    true,
		Secure:   session.PreferSecure,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   7 * 24 * time.Hour,
		Store: (&store.SQL{
			DB: db,
		}).GeneratePostgreSQLStatement("sessions", true),
	}))
	app.Server().UseFunc(pgctx.Middleware(db))

	mux := httpmux.New()
	todoapp.Mount(mux, app)

	app.Handler(mux)
	app.Address(":8080")

	slog.Info("starting server at :8080")

	err = app.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
