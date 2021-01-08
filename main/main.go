package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/tkircsi/urlshort"
)

func main() {
	yamlFile := flag.String("yaml", "paths.yaml", "read the paths and urls from the yaml file")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/quiz":  "https://github.com/tkircsi/quiz",
		"/go-ex": "https://gobyexample.com/command-line-flags",
		"/doc":   "https://godoc.org",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// 	yaml := `
	// - path: /quiz
	//   url: "https://github.com/tkircsi/quiz"
	// - path: /go-ex
	//   url: "https://gobyexample.com/command-line-flags"
	// `

	// 	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	yamlHandler, err := urlshort.YAMLHandlerFile(*yamlFile, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :5000")
	http.ListenAndServe(":5000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
