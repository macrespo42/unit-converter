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

func parseTemplates(templatePath string, r *http.Request) (formResult, *template.Template, error) {
	t, err := template.ParseFiles("./templates/base.html", templatePath)
	if err != nil {
		return formResult{}, t, err
	}

	rawLength := r.FormValue("length")
	from := r.FormValue("from")
	to := r.FormValue("to")

	data := formResult{
		Length: rawLength,
		From:   from,
		To:     to,
	}
	return data, t, nil
}

func lengthHandler(w http.ResponseWriter, r *http.Request) {
	data, t, err := parseTemplates("./templates/length.html", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if data.Length == "" {
		w.Header().Add("Content-Type", "text/html")
		err = t.Execute(w, nil)
		return
	}

	result, err := convert.ConvertLength(data.Length, data.From, data.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Result = result
	w.Header().Add("Content-Type", "text/html")
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func weightHandler(w http.ResponseWriter, r *http.Request) {
	data, t, err := parseTemplates("./templates/weight.html", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if data.Length == "" {
		w.Header().Add("Content-Type", "text/html")
		err = t.Execute(w, nil)
		return
	}

	result, err := convert.ConvertWeight(data.Length, data.From, data.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Result = result
	w.Header().Add("Content-Type", "text/html")
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	data, t, err := parseTemplates("./templates/temperature.html", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if data.Length == "" {
		w.Header().Add("Content-Type", "text/html")
		err = t.Execute(w, nil)
		return
	}

	result, err := convert.ConvertTemperature(data.Length, data.From, data.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Result = result
	w.Header().Add("Content-Type", "text/html")
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", lengthHandler)
	http.HandleFunc("/length", lengthHandler)
	http.HandleFunc("/weight", weightHandler)
	http.HandleFunc("/temperature", temperatureHandler)
	http.ListenAndServe(":8080", nil)
}
