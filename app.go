package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainPage)

	port := ":9999"
	println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}

}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

// comands for building on Ubuntu:
// export GOPATH=$HOME/Стільниця/programming/Golang
// go install go-tour
// $GOPATH/bin/go-tour

// go run app.go
