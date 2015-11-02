package commands

import (
	"go/build"
	"log"

	"github.com/bunniesandbeatings/go-flavor-parser/parser"
	"github.com/davecgh/go-spew/spew"
	"os"
	"github.com/cloudfoundry/gofileutils/fileutils"
	"fmt"
)

type ParseContext struct {
	ImportSpec string
	GOPATH     string
	GOROOT     string
	OutPath    string
	out        *os.File
}

func createBuildContext(parseContext *ParseContext) build.Context {
	buildContext := build.Default

	buildContext.GOPATH = parseContext.GOPATH
	buildContext.GOROOT = parseContext.GOROOT

	if parseContext.OutPath == "" {
		parseContext.out = os.Stdout;
	} else {
		out, e := fileutils.Create(parseContext.OutPath)
		if e != nil {
			log.Fatal(fmt.Printf("Could not create %s, failed with %s", parseContext.OutPath, e));
		}
		parseContext.out = out;
	}

	return buildContext
}

func ParsePackages(parseContext *ParseContext) {
	buildContext := createBuildContext(parseContext)

	parser := parser.NewParser(buildContext)
	parser.ParseImportSpec(parseContext.ImportSpec)

	log.Println(spew.Sdump(parser.GetArchitecture()))

	parseContext.out.Close();

}
