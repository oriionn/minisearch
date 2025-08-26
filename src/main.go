package main

import (
	"embed"
	"fmt"
	"minisearch/src/pages"
	"minisearch/src/utils"
	"net/http"
)

//go:embed public/*
var public embed.FS

func main() {
	// Classic routes
	http.HandleFunc("/", pages.Index)
	http.HandleFunc("/search", pages.Search)

	// Serve public files
	if (utils.DevMode()) {
		fileServer := http.FileServer(http.Dir("src/public"))
		http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	} else {
		fileServer := http.FileServer(http.FS(public))
		http.Handle("/public/", fileServer)
	}

	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
