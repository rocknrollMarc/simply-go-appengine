package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", indexPage)
}

// START FUNCTIONS FOR INDEX PAGE

var indexTemplate = template.Must(template.New("index.html").ParseFiles(
	"views/index.html",
))

type IndexPageVariables struct {
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	if err := indexTemplate.Execute(w, IndexPageVariables{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// END FUNCTIONS FOR INDEX PAGE
