package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ajstarks/mdtopdf"
)

var input = flag.String("i", "", "input markdown file (default is standard input")
var output = flag.String("o", "", "output PDF file (required)")
var tracefile = flag.String("t", "", "trace file")

func die(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *output == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// get text for PDF
	var content []byte
	var err error
	if *input == "" {
		content, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			die(err)
		}
	} else {
		content, err = os.ReadFile(*input)
		if err != nil {
			die(err)
		}
	}
	pf := mdtopdf.NewPdfRenderer("", "", *output, *tracefile)
	err = pf.Process(content)
	if err != nil {
		die(err)
	}
}
