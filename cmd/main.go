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
	todos, err := queries.ListTodos(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(todos)

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
	http.HandleFunc("/button-click", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path.Join(pwd, "layouts/response.html"))
		if err != nil {
			log.Fatal(err.Error())
		}
		t.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
