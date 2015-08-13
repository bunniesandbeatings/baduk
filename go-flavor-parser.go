package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"

	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"runtime/debug"

	"github.com/bunniesandbeatings/gotool"
	"github.com/davecgh/go-spew/spew"

	"github.com/bunniesandbeatings/go-flavor-parser/parser"
)

func usage() {
	appName := os.Args[0]
	log.Printf("%s usage:\n", appName)
	log.Printf("\t%s [flags] packages # see 'go help packages'\n", appName)
	log.Printf("Flags:\n")
	flag.PrintDefaults()
}

func importPaths(buildContext build.Context, importSpec string) []string {
	gotool.SetContext(buildContext)
	return gotool.ImportPaths([]string{importSpec})
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			debug.PrintStack()
			os.Exit(1)
		}
	}()

	commandContext := contexts.CreateCommandContext(usage)
	buildContext := contexts.CreateBuildContext(commandContext)

	importPaths := importPaths(buildContext, commandContext.ImportSpec)

	parser := parser.NewParser(buildContext)

	for _, importPath := range importPaths {
		parser.ParseImport(importPath)
	}

	fmt.Println(">>>> DEBUG: ast")
	spew.Dump(parser.Arch)
	fmt.Println("<<<< DEBUG: ast\n\n")

	//	datafile := datafiles.NewDataFile("com.bunniesandbeatings.go-flavor")
	//
	//	allFiles.ForEach(func(importPath string, files Any) {
	//		module := datafiles.Module{
	//			Name: importPath,
	//			Id: importPath,
	//			Type: "package",
	//		}
	//
	//		datafile.Modules = append(datafile.Modules, module)
	//		log.Printf("%s:", importPath)
	//
	//		for _, file := range files.([]*ast.File) {
	//			log.Printf("\t%#v:", *file)
	//		}
	//
	//	})

	//	ioutil.WriteFile(commandContext.OutputPath, datafile.ToXML(), 0644)
	//
	//	fmt.Printf("Output written to '%s'\n", commandContext.OutputPath)
	//	_ = packages.Packages(allFiles)
}
