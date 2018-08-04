package app

import (
	"strings"

	"github.com/acoshift/hime"

	"github.com/acoshift/todo-hime/repository"
)

func createGetHandler(ctx *hime.Context) error {
	return ctx.View("create", page(ctx))
}

func createPostHandler(ctx *hime.Context) error {
	f := getSession(ctx).Flash()

	content := strings.TrimSpace(ctx.PostFormValue("content"))

	if content == "" {
		f.Add("Errors", "content required")
		return ctx.RedirectToGet()
	}

	_, err := repository.CreateTodo(db, &repository.CreateTodoModel{
		Content: content,
	})
	if err != nil {
		return err
	}

	f.Set("Success", "todo successfully created")

	return ctx.RedirectTo("index")
}

func removePostHandler(ctx *hime.Context) error {
	f := getSession(ctx).Flash()

	todoID := strings.TrimSpace(ctx.PostFormValue("id"))

	if todoID == "" {
		f.Set("Errors", "id required")
		return ctx.RedirectTo("index")
	}

	err := repository.RemoveTodo(db, todoID)
	if err != nil {
		return err
	}

	f.Set("Success", "Todo successfully removed")

	return ctx.RedirectTo("index")
}

func donePostHandler(ctx *hime.Context) error {
	f := getSession(ctx).Flash()

	todoID := strings.TrimSpace(ctx.PostFormValue("id"))

	if todoID == "" {
		f.Set("Errors", "id required")
		return ctx.RedirectTo("index")
	}

	err := repository.SetTodoDone(db, todoID, true)
	if err != nil {
		return err
	}

	f.Set("Success", "Todo successfully mark as done")

	return ctx.RedirectTo("index")
}
