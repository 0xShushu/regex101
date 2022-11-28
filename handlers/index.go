package handlers

import (
	"net/http"
	"html/template"
	"regexp"
)

type Page struct {
	Result string
}

//render to index.html
func Index(w http.ResponseWriter, r *http.Request) {
	//make a new template
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//make a new page
	page := new(Page)

	//get queries
	regex := r.URL.Query().Get("regex")
	text := r.URL.Query().Get("text")

	//check params
	if regex != "" && text != "" {
		match, err := regexp.Match(regex, []byte(text))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if match {
			page.Result = "Regex match text"
		} else {
			page.Result = "Regex doesn't match text"
		}
	}

	//execute template
	if err := t.Execute(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}