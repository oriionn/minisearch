package main

import (
	"embed"
	"fmt"
	"minisearch/src/pages"
	"minisearch/src/utils"
	"net/http"
	"os"
	"slices"
	"strconv"
)

//go:embed public/*
var public embed.FS

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Printf("A request was made to %s\n", r.URL.Path)

	})
}

func main() {
	mux := http.NewServeMux()

	// Classic routes
	mux.HandleFunc("/", pages.Index)
	mux.HandleFunc("/search", pages.Search)

	// Serve public files
	if (utils.DevMode()) {
		fileServer := http.FileServer(http.Dir("src/public"))
		mux.Handle("/public/", http.StripPrefix("/public/", fileServer))
	} else {
		fileServer := http.FileServer(http.FS(public))
		mux.Handle("/public/", fileServer)
	}

	port := 3000
	if slices.Contains(os.Args, "--port") || slices.Contains(os.Args, "-p") {
		i := slices.Index(os.Args, "--port")
		if i == -1 {
			i = slices.Index(os.Args, "-p")
		}

		if i == len(os.Args) - 1 {
			fmt.Println("You have not specified a port.")
			return
		}

		num, err := strconv.Atoi(os.Args[i + 1])
		if err != nil {
			fmt.Println("You have not specified a port.")
			return
		}

		port = num
	}

	handler := LogMiddleware(mux)

	fmt.Printf("The HTTP server now runs on 0.0.0.0:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
