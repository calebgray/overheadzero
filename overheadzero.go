package main

import (
	"fmt"
	"github.com/calebgray/golibs/arguments"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
)

// run a development server.
func run(args []string) int {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/shutdown" {
			os.Exit(0)
		}

		fmt.Fprintf(w, "Welcome to %q", html.EscapeString(r.URL.Path))
	})

	var err error
	if len(args) == 0 {
		println("Starting a Server (:5000)")
		err = http.ListenAndServe(":5000", nil)
	} else {
		println("Starting a Server (", strings.Join(args, " "), ")")
		err = http.ListenAndServe(args[0], nil)
	}
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return 0
}

// main processes program commands and arguments.
func main() {
	arguments.AddCommand("run", "Run a development server.", run)
	arguments.Run()

	os.Exit(0)
}
