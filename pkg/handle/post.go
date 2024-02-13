package handle

// TODO: posted data will be handled here

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

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
	response := art.RunWebArtSingleLine(systemInput) // starts the decoder App and returns the happy path
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

func PostMulti(w http.ResponseWriter, r *http.Request) {
	var fileData models.ResponseBody
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading file content", http.StatusInternalServerError)
		}
		lines := strings.Split(string(content), "\n")
		fileData.FileContent = lines
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Fprintf(w, "File uploaded successfully")
	}

	pageHtml, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(fileData.FileContent)
	err = pageHtml.Execute(w, fileData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
