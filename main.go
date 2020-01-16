// Package documention here
package main

import (
	"fmt"
	"log"
	"net/http"

	// _ "net/http/pprof"
	"regexp"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	http.HandleFunc("/regexp/", handlerRegex)
	// http.HandleFunc("/", handlerRoot)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// func handlerRoot(w http.ResponseWriter, r *http.Request) {
// 	_, err := fmt.Fprintf(w, "Hello, world!")
// 	if err != nil {
// 		log.Fatalf("could not write to response %s", err)
// 	}
// }

var re = regexp.MustCompile("^(.+)@golang.org$")

func handlerRegex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")
	path := r.URL.Path
	match := re.FindAllStringSubmatch(path, -1)
	if match != nil {
		fmt.Fprintf(w, "Hello gopher %s", match[0][1])
		return
	}
	fmt.Fprint(w, "Hello, world!!")
}
