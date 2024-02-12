package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	SubmittedText string
}

func handler(w http.ResponseWriter, r *http.Request) {
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

func submitHandler(w http.ResponseWriter, r *http.Request) {
	textInput := r.FormValue("textInput")
	data := PageData{
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

func main() {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/decoder", submitHandler)
	fmt.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
