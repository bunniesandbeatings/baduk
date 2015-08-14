package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"
	"github.com/bunniesandbeatings/go-flavor-parser/parser"

	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/davecgh/go-spew/spew"

)

func usage() {
	appName := os.Args[0]
	log.Printf("%s usage:\n", appName)
	log.Printf("\t%s [flags] packages # see 'go help packages'\n", appName)
	log.Printf("Flags:\n")
	flag.PrintDefaults()
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

	parser := parser.NewParser(buildContext)

	parser.ParseImportSpec(commandContext.ImportSpec)

	fmt.Println(">>>> DEBUG: ast")
	spew.Dump(parser.GetArchitecture())
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
