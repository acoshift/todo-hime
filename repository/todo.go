package repository

import "github.com/acoshift/todo-hime/entity"

// CreateTodoModel is the model for CreateTodo
type CreateTodoModel struct {
	Content string
	Done    bool
}

// CreateTodo creates new todo
func CreateTodo(q Queryer, x *CreateTodoModel) (todoID string, err error) {
	err = q.QueryRow(`
		insert into todos
			(content, done)
		values
			($1, $2)
		returning id
	`, x.Content, x.Done).Scan(&todoID)
	return
}

// SetTodoDone sets todo done state
func SetTodoDone(q Queryer, todoID string, done bool) error {
	_, err := q.Exec(`
		update todos
		set
			done = $2
		where id = $1
	`, todoID, done)
	return err
}

// RemoveTodo removes a todo
func RemoveTodo(q Queryer, todoID string) error {
	_, err := q.Exec(`
		delete from todos
		where id = $1
	`, todoID)
	return err
}

// ListTodos lists all todos
func ListTodos(q Queryer) ([]*entity.Todo, error) {
	rows, err := q.Query(`
		select
			id, content, done, created_at
		from todos
		order by created_at desc
	`)
	if err != nil {
		return nil, err
	}

	xs := make([]*entity.Todo, 0)
	for rows.Next() {
		var x entity.Todo
		err = rows.Scan(
			&x.ID, &x.Content, &x.Done, &x.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		xs = append(xs, &x)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return xs, nil
}
