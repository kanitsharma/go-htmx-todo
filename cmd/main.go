package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
	"kanitsharma.dev/go-htmx-todo/todos"
)

//go:embed db/schema.sql
var ddl string

func main() {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err.Error())
	}

	queries := todos.New(db)

	data := map[string]interface{}{
		"Title": "testpage",
	}
	pwd, _ := os.Getwd()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join(pwd, "layouts/page.html"))
		if err != nil {
			log.Fatal(err.Error())
		}
		t.Execute(w, data)
	})
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos, err := queries.ListTodos(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
		data := map[string]interface{}{
			"Todos": todos,
		}
		log.Print(todos)
		t, err := template.ParseFiles(path.Join(pwd, "layouts/todos.html"))
		if err != nil {
			log.Fatal(err.Error())
		}
		t.Execute(w, data)
	})
	http.HandleFunc("/add-todos", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.Print(r.Form)
		todo, err := queries.CreateTodo(ctx, todos.CreateTodoParams{
			Name: r.Form["name"][0],
		})
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Print(todo)

		data := map[string]interface{}{
			"Todos": append(make([]todos.Todo, 0), todo),
		}
		t, err := template.ParseFiles(path.Join(pwd, "layouts/todos.html"))
		if err != nil {
			log.Fatal(err.Error())
		}
		t.Execute(w, data)
	})
	http.HandleFunc("/delete-todo", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		err := queries.DeleteTodo(ctx, name)
		if err != nil {
			log.Fatal(err.Error())
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
