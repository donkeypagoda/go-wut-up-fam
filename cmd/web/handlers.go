package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	files := []string{
		"./ui/html/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "home", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	json, err := os.ReadFile("./internal/api.json")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Fprintf(w, "Here's some really important data for you: %s", json)
}
