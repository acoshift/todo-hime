package app

import (
	"github.com/acoshift/hime"

	"github.com/acoshift/todo-hime/repository"
)

func indexGetHandler(ctx *hime.Context) error {
	list, err := repository.ListTodos(db)
	if err != nil {
		return err
	}

	p := page(ctx)
	p["List"] = list
	return ctx.View("index", p)
}
