package app

import (
	"database/sql"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/acoshift/header"
	"github.com/acoshift/hime"
	"github.com/acoshift/httprouter"
	"github.com/acoshift/middleware"
	"github.com/acoshift/session"
	"github.com/acoshift/webstatic"
)

// app's global vars
var (
	db  *sql.DB
	loc *time.Location
)

// New creates new handler
func New(app *hime.App, c Config) http.Handler {
	db = c.DB
	loc = c.Location

	app.
		Routes(hime.Routes{
			"index":  "/",
			"create": "/create",
			"done":   "/done",
			"remove": "/remove",
		})

	mux := http.NewServeMux()
	mux.Handle("/-/", http.StripPrefix("/-", webstatic.New(webstatic.Config{
		Dir:          "assets",
		CacheControl: "public, max-age=31536000",
	})))

	r := httprouter.New()
	r.HandleMethodNotAllowed = false
	r.HandleOptions = false
	r.NotFound = hime.Handler(notFoundHandler)

	r.Get(app.Route("index"), hime.Handler(indexGetHandler))
	r.Get(app.Route("create"), hime.Handler(createGetHandler))
	r.Post(app.Route("create"), hime.Handler(createPostHandler))
	r.Post(app.Route("done"), hime.Handler(donePostHandler))
	r.Post(app.Route("remove"), hime.Handler(removePostHandler))

	mux.Handle("/", middleware.Chain(
		session.Middleware(session.Config{
			Rolling:  true,
			Proxy:    true,
			Secure:   session.PreferSecure,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
			HTTPOnly: true,
			MaxAge:   7 * 24 * time.Hour,
			Store:    c.SessionStore,
		}),
	)(r))

	return middleware.Chain(
		panicRecovery,
		noCORS,
	)(mux)
}

func defaultCacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(header.CacheControl, "no-cache, no-store, must-revalidate")
		h.ServeHTTP(w, r)
	})
}

func noCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			http.Error(w, "Forbidded", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func panicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println(err)
				debug.PrintStack()
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func notFoundHandler(ctx *hime.Context) error {
	return ctx.RedirectTo("index")
}
