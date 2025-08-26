package main

import (
	"embed"
	"fmt"
	"meta-searchengine/src/pages"
	"net/http"
)

//go:embed public/*
var public embed.FS

func main() {
	// Classic routes
	http.HandleFunc("/", pages.Index)
	http.HandleFunc("/search", pages.Search)

	// Serve public files
	fileServer := http.FileServer(http.FS(public))
	http.Handle("/public/", fileServer)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
