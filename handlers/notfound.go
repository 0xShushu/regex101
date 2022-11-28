package handlers

import (
	"net/http"
	"html/template"
)

//render to 404.html
func NotFound(w http.ResponseWriter, r *http.Request) {
	//make new template
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//execute template
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}