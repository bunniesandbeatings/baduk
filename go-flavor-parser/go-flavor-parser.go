package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/files"
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"
	"github.com/bunniesandbeatings/go-flavor-parser/structure101/datafiles"

	//	"github.com/bunniesandbeatings/go-flavor-parser/packages"
	. "github.com/mndrix/ps"

	"go/ast"
	
	"io/ioutil"
	"flag"
	"log"
	"fmt"
	"os"
)

func usage() {
	appName := os.Args[0]
	log.Printf("%s usage:\n", appName)
	log.Printf("\t%s [flags] packages # see 'go help packages'\n", appName)
	log.Printf("Flags:\n")
	flag.PrintDefaults()
}

func main() {
	commandContext := contexts.CreateCommandContext(usage)
	buildContext := contexts.CreateBuildContext(commandContext)

	allFiles := files.Files(buildContext, commandContext.ImportSpec)

	//	allFiles.ForEach(func(key string, value ps.Any) {
	//		log.Printf("%s: %#v", key, value)
	//	})

	datafile := datafiles.NewDataFile("com.bunniesandbeatings.go-flavor")

	allFiles.ForEach(func(importPath string, files Any) {
		module := datafiles.Module{
			Name: importPath,
			Id: importPath,
			Type: "package",
		}
		
		datafile.Modules = append(datafile.Modules, module)
		log.Printf("%s:", importPath)

		for _, file := range files.([]*ast.File) {
			log.Printf("\t%#v:", *file)
		}

	})

	ioutil.WriteFile(commandContext.OutputPath, datafile.ToXML(), 0644)

	fmt.Printf("Output written to '%s'\n", commandContext.OutputPath)
	//	_ = packages.Packages(allFiles)
}
