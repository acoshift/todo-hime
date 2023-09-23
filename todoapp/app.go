package todoapp

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/acoshift/methodmux"
	"github.com/acoshift/webstatic"
	"github.com/moonrhythm/hime"
	"github.com/moonrhythm/httpmux"

	"github.com/acoshift/todo-hime/template"
)

func Mount(m *httpmux.Mux, app *hime.App) http.Handler {
	t := app.Template()
	t.FS(template.FS)
	t.Root("root")
	t.Minify()
	t.Preload("layout.tmpl")
	t.ParseComponentFile("todo-table-body", "component/todo-table-body.tmpl")
	t.ParseFiles("index", "index.tmpl")
	t.ParseFiles("create", "create.tmpl")

	app.Routes(hime.Routes{
		"index":  "/",
		"create": "/create",
		"done":   "/done",
		"remove": "/remove",
		"list":   "/list",
	})

	m = m.Group("/", panicRecovery, noCORS)

	m.Handle("/-/", http.StripPrefix("/-", webstatic.New(webstatic.Config{
		Dir:          "assets",
		CacheControl: "public, max-age=31536000",
	})))

	m.Handle("/", methodmux.Get(hime.Handler(getIndex)))
	m.Handle("/list", methodmux.Get(hime.Handler(getList)))
	m.Handle("/create", methodmux.GetPost(
		hime.Handler(getCreate),
		hime.Handler(postCreate),
	))
	m.Handle("/done", methodmux.Post(hime.Handler(postDone)))
	m.Handle("/remove", methodmux.Post(hime.Handler(postRemove)))

	return m
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
