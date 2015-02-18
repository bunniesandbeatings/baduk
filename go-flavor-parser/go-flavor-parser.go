package main

import (
	"github.com/bunniesandbeatings/gotool"
	"github.com/bunniesandbeatings/go-flavor-parser/files"

	"github.com/tonnerre/golang-pretty"

	"flag"
	"fmt"
	"os"
	"go/build"
)

var (
	goPath string
	goRoot string
	outputPath string
	importSpec string
)

func usage() {
	appName := os.Args[0]
	fmt.Fprintf(os.Stderr, "%s usage:\n", appName)
	fmt.Fprintf(os.Stderr, "\t%s [flags] packages # see 'go help packages'\n", appName)
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	handleCommandLine()

	fmt.Printf("Ran with arguments: %s\n", os.Args)

	buildContext := buildContext()
	fmt.Printf("\n*** Context ***\n%# v\n\n", pretty.Formatter(buildContext))

	gotool.SetContext(buildContext)
	projectImportPaths := gotool.ImportPaths([]string{importSpec})

	_, allFiles := files.GetFiles(buildContext, nil, projectImportPaths)

	//	datafile = Parse(buildContext, projectImportPaths)

	//	ioutil.WriteFile(outputPath, parser.DataFileXML(), 0644)

	//	fmt.Printf("Output written to '%s'\n", outputPath)
	fmt.Printf("\n*** all packages ***\n%# v\n\n", pretty.Formatter(allFiles))

}

func handleCommandLine() {
	flag.StringVar(&goPath, "gopath", os.Getenv("GOPATH"), "allows you to choose a different GOPATH to use during analysis")

	flag.StringVar(&goRoot, "goroot", os.Getenv("GOROOT"), "allows you to choose a different GOROOT to use during analysis")
	flag.StringVar(&outputPath, "output", "./output.xml", "where to output the result of the analysis")

	flag.Usage = usage
	flag.Parse()

	importSpec = flag.Arg(0)
}

func buildContext() build.Context {
	buildContext := build.Default

	if goPath != "" {
		buildContext.GOPATH = goPath
	}

	if goRoot != "" {
		buildContext.GOROOT = goRoot
	}

	return buildContext
}