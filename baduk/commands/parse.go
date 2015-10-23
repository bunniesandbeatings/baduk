package commands

import (
	"go/build"
	"log"

	"github.com/bunniesandbeatings/go-flavor-parser/parser"
	"github.com/davecgh/go-spew/spew"
)

type ParseContext struct {
	ImportSpec string
	GOPATH     string
	GOROOT     string
	//	Out        *os.File
}

func createBuildContext(parseContext *ParseContext) build.Context {
	buildContext := build.Default

	buildContext.GOPATH = parseContext.GOPATH
	buildContext.GOROOT = parseContext.GOROOT

	return buildContext
}

func ParsePackages(parseContext *ParseContext) {
	buildContext := createBuildContext(parseContext)

	parser := parser.NewParser(buildContext)

	parser.ParseImportSpec(parseContext.ImportSpec)

	log.Println(spew.Sdump(parser.GetArchitecture()))
}
