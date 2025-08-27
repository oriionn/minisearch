package main

import (
	"embed"
	"flag"
	"fmt"
	"minisearch/src/pages"
	"net/http"
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

	var (
		port = flag.Int("port", 3000, "The port where the web server listens")
		dev = flag.Bool("dev", false, "Enable the dev mode")
	)

	flag.IntVar(port, "p", 3000, "The port where the web server listens")
	flag.BoolVar(dev, "d", false, "Enable the dev mode")

	flag.Parse()

	// Classic routes
	mux.HandleFunc("/", pages.Index)
	mux.HandleFunc("/search", pages.Search)

	// Serve public files
	if *dev {
		fileServer := http.FileServer(http.Dir("src/public"))
		mux.Handle("/public/", http.StripPrefix("/public/", fileServer))
	} else {
		fileServer := http.FileServer(http.FS(public))
		mux.Handle("/public/", fileServer)
	}

	handler := LogMiddleware(mux)

	fmt.Printf("The HTTP server now runs on 0.0.0.0:%d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handler)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
