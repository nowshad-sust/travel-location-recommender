package controller

import (
	"path"
	"net/http"
	"html/template"

	"github.com/julienschmidt/httprouter"
)


func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	fp := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, false); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}