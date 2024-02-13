package handle

// TODO: posted data will be handled here

import (
	"html/template"
	"net/http"

	"gitea.kood.tech/hannessoosaar/art/pkg/art"
	"gitea.kood.tech/hannessoosaar/art/pkg/models"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	textInput := r.FormValue("textInput")
	systemInput := []string{"placeHolder", "-e", textInput}
	art.InitializeAndRun(systemInput)
	data := models.PageData{ 		// creates an instance of PageData
		SubmittedText: textInput,
	}
	pageHtml, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pageHtml.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}