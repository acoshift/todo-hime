package app

import (
	"github.com/acoshift/hime"

	"github.com/acoshift/todo-hime/repository"
)

func indexGetHandler(ctx *hime.Context) hime.Result {
	list, err := repository.ListTodos(db)
	must(err)

	p := page(ctx)
	p["List"] = list
	return ctx.View("index", p)
}
