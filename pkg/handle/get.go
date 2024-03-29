package handle

import (
	"html/template"
	"net/http"
)

func LoadIndex(w http.ResponseWriter, r *http.Request) {
	pageHtml, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pageHtml.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
