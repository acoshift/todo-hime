# todo-hime

Hime Example for Todo app

## Running

1. Prepare database

    ```bash
    $ psql -U postgres -c "create database todo_hime"
    $ psql -U postgres -d todo_hime -f schema.sql
    ```

2. Run server

    ```bash
    $ go run .
    ```
