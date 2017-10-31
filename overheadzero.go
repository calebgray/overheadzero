package main

import (
	"fmt"
	"github.com/calebgray/golibs/arguments"
	"html"
	"log"
	"net/http"
	"os"
)

// run a development server.
func run(args []string) int {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/shutdown" {
			os.Exit(0)
		}

		fmt.Fprintf(w, "Welcome to %q", html.EscapeString(r.URL.Path))
	})

	var addr string
	if len(args) == 0 {
		addr = ":5000"
	} else {
		addr = args[0]
	}
	println("Starting a Server (", addr, ")")
	err := http.ListenAndServe(addr, nil)
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
