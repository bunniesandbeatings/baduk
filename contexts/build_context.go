package contexts

import (
	"github.com/tonnerre/golang-pretty"

	"go/build"
	"log"
)

func CreateBuildContext(commandContext CommandContext) build.Context {
	buildContext := build.Default

	if commandContext.GoPath != "" {
		buildContext.GOPATH = commandContext.GoPath
	}

	if commandContext.GoRoot != "" {
		buildContext.GOROOT = commandContext.GoRoot
	}

	log.Printf("\n*** Build Context ***\n%# v\n\n", pretty.Formatter(buildContext))

	return buildContext
}
