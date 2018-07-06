/*
Package main, this one, is for the parser
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jdholdren/airblog/pkg/parsing"
)

func main() {
	// Get the first command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage is: airblog [file]")
		return
	}
	filename := os.Args[1]

	// Get a reader for that file
	r, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Make the parser
	p := parsing.NewParser(r)

	p.Run()
}
