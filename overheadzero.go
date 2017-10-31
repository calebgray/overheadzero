package main

import (
	"fmt"
	"github.com/calebgray/golibs/arguments"
	"html"
	"log"
	"net/http"
	"os"
)

// run a development server
func run(args []string) int {
	// read the arguments
	var addr string
	if len(args) == 0 {
		addr = ":5000"
	} else {
		addr = args[0]
	}

	// start the http server
	println("Starting a Server (", addr, ")")
	err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/shutdown" {
			os.Exit(0)
		}

		// todo: serve static/dynamic/compiled content
		fmt.Fprintf(w, "Welcome to %q", html.EscapeString(r.URL.Path))
	}))
	if err != nil {
		log.Fatal("Error: ", err)
		return 1
	}

	return 0
}

// compile the sources in a directory
func compile(args []string) int {
	var sources []string
	if len(args) == 0 {
		// todo: find all the files that can be compiled
	} else {
		sources = args
	}

	for _, source := range sources {
		println(source)
	}

	return 0
}

// main processes program commands and arguments
func main() {
	arguments.AddCommand("run", "Run a development server", run)
	arguments.AddCommand("compile", "Compile the sources in a directory", compile)
	arguments.Run()

	os.Exit(0)
}
