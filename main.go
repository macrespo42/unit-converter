package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/macrespo42/unit-converter/internal/convert"
)

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

	length, err := strconv.ParseFloat(rawLength, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	milimeters, err := convert.ToMilimeter(length, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := convert.MilimeterTo(milimeters, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Length float64
		From   string
		To     string
		Result float64
	}{
		Length: length,
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
