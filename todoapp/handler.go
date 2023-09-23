package todoapp

import (
	"strings"

	"github.com/moonrhythm/hime"
)

func getIndex(ctx *hime.Context) error {
	if ctx.URL.Path != "/" {
		return ctx.RedirectTo("index")
	}

	list, err := ListTodo(ctx)
	if err != nil {
		return err
	}

	p := page(ctx)
	p["List"] = list
	return ctx.View("index", p)
}

func getList(ctx *hime.Context) error {
	// time.Sleep(200 * time.Millisecond) // simulate slow query

	list, err := ListTodo(ctx)
	if err != nil {
		return err
	}

	return ctx.Component("todo-table-body", list)
}

func getCreate(ctx *hime.Context) error {
	return ctx.View("create", page(ctx))
}

func postCreate(ctx *hime.Context) error {
	f := getSession(ctx).Flash()

	content := strings.TrimSpace(ctx.PostFormValue("content"))

	if content == "" {
		f.Add("Errors", "content required")
		return ctx.RedirectToGet()
	}

	_, err := CreateTodo(ctx, &CreateTodoParams{
		Content: content,
	})
	if err != nil {
		return err
	}

	f.Set("Success", "todo successfully created")

	return ctx.RedirectTo("index")
}

func postRemove(ctx *hime.Context) error {
	f := getSession(ctx).Flash()

	todoID := strings.TrimSpace(ctx.PostFormValue("id"))

	if todoID == "" {
		f.Set("Errors", "id required")
		return ctx.RedirectBackToGet()
	}

	err := RemoveTodo(ctx, todoID)
	if err != nil {
		return err
	}

	f.Set("Success", "Todo successfully removed")

	return ctx.RedirectBackToGet()
}

func postDone(ctx *hime.Context) error {
	f := getSession(ctx).Flash()

	todoID := strings.TrimSpace(ctx.PostFormValue("id"))

	if todoID == "" {
		f.Set("Errors", "id required")
		return ctx.RedirectBackToGet()
	}

	err := SetTodoDone(ctx, todoID, true)
	if err != nil {
		return err
	}

	f.Set("Success", "Todo successfully mark as done")

	return ctx.RedirectBackToGet()
}
