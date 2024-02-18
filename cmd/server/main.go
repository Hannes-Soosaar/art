package main

import (
	"fmt"
	"net/http"

	"gitea.kood.tech/hannessoosaar/art/conf"
	"gitea.kood.tech/hannessoosaar/art/pkg/handle"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handle.LoadIndex)
	http.HandleFunc("/decoder", handle.PostEncoded)
	http.HandleFunc("/encoder", handle.PostEncoded)
	http.HandleFunc("/reload", handle.LoadIndex)
	http.HandleFunc("/multi", handle.PostMulti)
	fmt.Printf("Server listening on :%s\n", conf.ServerPort)
	err := http.ListenAndServe(conf.ServerPort, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
