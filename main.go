package main

import (
	"html/template"
	"net/http"

	"github.com/macrespo42/unit-converter/internal/convert"
)

type formResult struct {
	Length string
	From   string
	To     string
	Result float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.html", "./templates/length.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rawLength := r.FormValue("length")
	from := r.FormValue("from")
	to := r.FormValue("to")

	if rawLength == "" {
		w.Header().Add("Content-Type", "text/html")
		err = t.Execute(w, nil)
		return
	}

	result, err := convert.ConvertLength(rawLength, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := formResult{
		Length: rawLength,
		From:   from,
		To:     to,
		Result: result,
	}

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
