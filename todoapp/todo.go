package todoapp

import (
	"context"
	"time"

	"github.com/acoshift/pgsql"
	"github.com/acoshift/pgsql/pgctx"
)

type Todo struct {
	ID        string
	Content   string
	Done      bool
	CreatedAt time.Time
}

type CreateTodoParams struct {
	Content string
	Done    bool
}

func CreateTodo(ctx context.Context, r *CreateTodoParams) (string, error) {
	var id string
	err := pgctx.QueryRow(ctx, `
		insert into todos (content, done)
		values ($1, $2)
		returning id
	`, r.Content, r.Done).Scan(&id)
	return id, err
}

func SetTodoDone(ctx context.Context, todoID string, done bool) error {
	_, err := pgctx.Exec(ctx, `
		update todos
		set done = $2
		where id = $1
	`, todoID, done)
	return err
}

func RemoveTodo(ctx context.Context, todoID string) error {
	_, err := pgctx.Exec(ctx, `
		delete from todos
		where id = $1
	`, todoID)
	return err
}

func ListTodo(ctx context.Context) ([]*Todo, error) {
	var xs []*Todo
	err := pgctx.Iter(ctx, func(scan pgsql.Scanner) error {
		var x Todo
		err := scan(&x.ID, &x.Content, &x.Done, &x.CreatedAt)
		if err != nil {
			return err
		}
		xs = append(xs, &x)
		return nil
	}, `
		select id, content, done, created_at
		from todos
		order by created_at desc
	`)
	if err != nil {
		return nil, err
	}
	return xs, nil
}
