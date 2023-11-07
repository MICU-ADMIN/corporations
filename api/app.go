package main

import (
	"embed"
	"html/template"
    	"net/http"
	"os"
)

//go:embed templates/*
var templates embed.FS
var t = template.Must(template.ParseFS(templates, "templates/*"))

//go:embed static/*
var assets embed.FS


func main() {
	port := "3000"
	if(os.Getenv("PORT") != "") {
		port = os.Getenv("PORT")
	}

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        t.Execute(w, "")
    })

    fs := http.FileServer(http.FS(assets))
    http.Handle("/static/", fs)

    http.ListenAndServe(":" + port, nil)
}
