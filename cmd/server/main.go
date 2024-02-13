package main

import (
	"fmt"
	"net/http"

	"gitea.kood.tech/hannessoosaar/art/conf"
	"gitea.kood.tech/hannessoosaar/art/pkg/handle"
)

func main() {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // sets the static folder for css
	http.HandleFunc("/", handle.Handler)
	http.HandleFunc("/decoder", handle.SubmitHandler)
	fmt.Printf("Server listening on :%s\n", conf.ServerPort)
	err := http.ListenAndServe(conf.ServerPort, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
