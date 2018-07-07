/*
Package main, this one, is for the parser
*/
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jdholdren/airblog/pkg/markdown"
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

	n := time.Now()
	p.Run(markdown.Initial)
	t := time.Now()

	fmt.Printf("Finished in %dms", (t.UnixNano()-n.UnixNano())/1000)
	fmt.Printf("\n")
}
