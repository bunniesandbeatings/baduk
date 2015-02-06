package main

import (
	. "github.com/bunniesandbeatings/go-flavor-parser/parser"

	"github.com/kisielk/gotool"

	"flag"
	"fmt"
	"os"
	"go/build"
)

var (
	gopath string
	showUsage bool
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

	if gopath != "" {
		buildContext := build.Default
		buildContext.GOPATH = gopath

		parser.Packages.BuildContext = buildContext
	}
	
	fmt.Printf("Using GOPATH='%s'\n", parser.Packages.BuildContext.GOPATH)
	
	parser.AddImportPaths(projectImportPaths)

	fmt.Println(string(parser.DataFileXML()))
}

func importPathsFromCommandLine() []string {
	flag.StringVar(&gopath, "gopath", os.Getenv("GOPATH"), "allows you to choose a different GOPATH to use during analysis")
	flag.Usage = usage
	flag.Parse()

	importSpec := []string{flag.Arg(0)}

	return gotool.ImportPaths(importSpec)
}

