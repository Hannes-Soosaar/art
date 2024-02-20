package handle

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

func PostSingle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing from data", http.StatusInternalServerError)
	}
	data := models.ResponseBody{ // creates an instance of PageData
		Options:       r.Form.Get("options"),
		SubmittedText: r.Form.Get("textInput"),
	}
	// fmt.Println(data.Options + "\n" + data.SubmittedText)
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
		fileData.Options = r.FormValue("options") // sets the option
		if fileData.Options == "" {
			fmt.Print("ERROR") // add proper error here!
		}
		fmt.Printf("\n type of Operations: %s", fileData.Options)
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
		var fileContent []string
		for _, line := range lines {
			fileContent = append(fileContent, line)
		}
		fileData = art.RunWebArtMultiLine(fileData.Options, fileContent)
		w.Header().Set("Content-Type", "html")
	}
	pageHtml, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pageHtml.Execute(w, fileData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
