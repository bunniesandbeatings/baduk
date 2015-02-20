package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/files"
	"github.com/bunniesandbeatings/go-flavor-parser/context"
	"github.com/bunniesandbeatings/go-flavor-parser/packages"

	"flag"
	"log"
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
	commandContext := context.CreateCommandContext(usage)
	buildContext := context.CreateBuildContext(commandContext)

	allFiles := files.GetFilesFromImportSpec(buildContext, commandContext.ImportSpec)

	_ = packages.GetPackagesFromFileList(allFiles)

}
