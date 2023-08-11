package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	data := map[string]interface{}{
		"Title": "testpage",
	}
	pwd, _ := os.Getwd()

	http.HandleFunc("/htmx", func(w http.ResponseWriter, r *http.Request) {
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
		t.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
