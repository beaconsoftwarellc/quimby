package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gitlab.com/beacon-software/gadget/stringutil"
	"gitlab.com/beacon-software/quimby/codegen/generator"
)

//go:generate go-embed template templates templates/templates.go

var (
	fileName = flag.String("f", "definition.yaml", "the path to the definition file.")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if stringutil.IsWhiteSpace(*fileName) {
		log.Printf("No definition specified.")
		Usage()
		os.Exit(1)
	}

	gen, err := generator.New(*fileName)
	if nil != err {
		log.Print(err)
		Usage()
		os.Exit(1)
	}

	err = gen.Run()
	if nil != err {
		log.Fatal(err)
	}
}
