package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.html", "./templates/length.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := struct{ Name string }{Name: "John"}

	w.Header().Add("Content-Type", "text/html")
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
