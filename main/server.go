package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("../static/")))

	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := "https://neeft-back.herokuapp.com"
	println("Server listen on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error : ", err)
	}

}
