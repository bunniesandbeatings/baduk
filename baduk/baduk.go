package main

import (
	"fmt"
	"os"

	"github.com/bunniesandbeatings/go-flavor-parser/baduk/commands"
	"github.com/jawher/mow.cli"
)

func main() {
	doc := `Intermediate Golang AST for Architectural Viz
Copyright (c) 2014 Rasheed Abdul-Aziz

License: MIT

Authors: Rasheed Abdul-Aziz, Glyn Normington`

	parse := cli.App("baduk", doc)

	gopath := parse.String(cli.StringOpt{
		Name:   "gopath",
		Value:  "",
		Desc:   "allows you to choose a different GOPATH to use during analysis",
		EnvVar: "GOPATH",
	})

	goroot := parse.String(cli.StringOpt{
		Name:   "goroot",
		Value:  "",
		Desc:   "allows you to choose a different GOROOT to use during analysis",
		EnvVar: "GOROOT",
	})

	packages := parse.StringArg("PACKAGES", "", "the package specification to parse")

	parse.Spec = "[OPTIONS] PACKAGES"

	parse.Version("v version", fmt.Sprintf("baduk %s", VERSION))

	parse.Action = func() {
		parseContext := &commands.ParseContext{
			GOROOT:     *goroot,
			GOPATH:     *gopath,
			ImportSpec: *packages,
		}

		commands.ParsePackages(parseContext)
	}

	parse.Run(os.Args)
}
