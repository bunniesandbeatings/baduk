package main

import (
	. "github.com/bunniesandbeatings/go-flavor-parser/parser"

	"github.com/kisielk/gotool"

	"flag"
	"fmt"
	"os"
	"go/build"
	"io/ioutil"
)

var (
	goPath string
	outputPath string
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

	if goPath != "" {
		buildContext := build.Default
		buildContext.GOPATH = goPath

		parser.Packages.BuildContext = buildContext
	}

	fmt.Printf("Using GOPATH='%s'\n", parser.Packages.BuildContext.GOPATH)
	
	parser.AddImportPaths(projectImportPaths)
	fmt.Printf("Ran with arguments: %s\n", os.Args)

	ioutil.WriteFile(outputPath, parser.DataFileXML(), 0644)
	
	fmt.Printf("Output written to '%s'\n", outputPath)
}

func importPathsFromCommandLine() []string {
	flag.StringVar(&goPath, "gopath", os.Getenv("GOPATH"), "allows you to choose a different GOPATH to use during analysis")
	flag.StringVar(&outputPath, "output", "./output.xml", "where to output the result of the analysis")

	flag.Usage = usage
	flag.Parse()

	importSpec := []string{flag.Arg(0)}

	return gotool.ImportPaths(importSpec)
}

