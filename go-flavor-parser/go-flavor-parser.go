package main

import (
	. "github.com/bunniesandbeatings/go-flavor-parser/parser"
	
	"github.com/kisielk/gotool"

	"flag"
	"fmt"
	"os"
)

func usage() {
	appName := os.Args[0]
	fmt.Fprintf(os.Stderr, "%s usage:\n", appName)
	fmt.Fprintf(os.Stderr, "\t%s [flags] packages # see 'go help packages'\n", appName)
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	projectImportPaths := importPathsFromCommandLine()
	
	parser := NewParser()
	parser.AddImports(projectImportPaths)

	fmt.Println(string(parser.DataFileXML()))
}

func importPathsFromCommandLine() []string {
	flag.Usage = usage
	flag.Parse()

	importSpec := []string{flag.Arg(0)}

	return gotool.ImportPaths(importSpec)
}

