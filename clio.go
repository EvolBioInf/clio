// The clio package implements functions to automate routine input and output operations in command-line programs.
package clio

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// FileAction takes as first, and possibly only, argument an io.Reader. In addition, it can take variadic arguments of type empty interface.
type FileAction func(io.Reader, ...interface{})

// ParseFiles iterates over a set of files and applies the same function to each one. The arguments of that function are variadic and have the empty interface type.
func ParseFiles(files []string, fn FileAction, args ...interface{}) {
	r := os.Stdin
	if len(files) == 0 {
		fn(r, args...)
	} else {
		for _, f := range files {
			r, err := os.Open(f)
			if err != nil {
				log.Fatalf("couldn't open %q\n", f)
			}
			fn(r, args...)
			r.Close()
		}
	}
}

// PrintInfo prints the name of a program, its version, date of compilation, author, her/his email address, and the program's license.
func PrintInfo(name, version, date, author, email, license string) {
	fmt.Printf("%s %s, %s\n", name, version, date)
	fmt.Printf("Author: %s, %s\n", author, email)
	fmt.Printf("License: %s\n", license)
}

// Usage sets the response to a request for help.
func Usage(usage, purpose, example string) {
	flag.Usage = func() {
		o := flag.CommandLine.Output()
		fmt.Fprintf(o, "Usage: %s\n%s\nExample: %s\n",
			usage, purpose, example)
		fmt.Fprintf(o, "Options:\n")
		flag.PrintDefaults()
	}
}
