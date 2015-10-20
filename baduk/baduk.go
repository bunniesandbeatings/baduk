package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"
	"github.com/bunniesandbeatings/go-flavor-parser/parser"

	"os"

	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	"log"
	"runtime/debug"
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

	app := cli.NewApp()
	app.Name = "baduk"
	app.Usage = "Intermediate Golang AST for Architectural Viz"

	app.Flags = contexts.CreateCommandContextFlags()

	app.Action = func(cliContext *cli.Context) {
		commandContext := contexts.CreateCommandContext(cliContext)
		buildContext := contexts.CreateBuildContext(commandContext)

		parser := parser.NewParser(buildContext)

		parser.ParseImportSpec(commandContext.ImportSpec)

		log.Println(spew.Sdump(parser.GetArchitecture()))
	}

	app.Run(os.Args)
}
