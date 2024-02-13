package handle

// TODO: posted data will be handled here

import (
	"fmt"
	"html/template"
	"net/http"

	// "gitea.kood.tech/hannessoosaar/art/pkg/art"
	"gitea.kood.tech/hannessoosaar/art/pkg/art"
	"gitea.kood.tech/hannessoosaar/art/pkg/models"
)

func PostEncoded(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing from data", http.StatusInternalServerError)
	}
	data := models.PageData{ // creates an instance of PageData
		Options:       r.Form.Get("options"),
		SubmittedText: r.Form.Get("textInput"),
	}
	fmt.Println(data.Options + "\n" + data.SubmittedText)

	systemInput := []string{data.Options, data.SubmittedText} 
	response := art.RunWebArtSingleLine(systemInput)          // starts the decoder App and returns the happy path
	pageHtml, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pageHtml.Execute(w, response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
