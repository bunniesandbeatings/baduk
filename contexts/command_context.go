package contexts

import (
	"log"
	"flag"
	"os"
)

type CommandContext struct {
	GoRoot     string
	GoPath     string
	OutputPath string
	ImportSpec string
}

func CreateCommandContext(usage func()) CommandContext {
	commandContext := CommandContext{}

	flag.StringVar(&commandContext.GoPath, "gopath", os.Getenv("GOPATH"), "allows you to choose a different GOPATH to use during analysis")
	flag.StringVar(&commandContext.GoRoot, "goroot", os.Getenv("GOROOT"), "allows you to choose a different GOROOT to use during analysis")

	flag.StringVar(&commandContext.OutputPath, "output", "./output.xml", "where to output the result of the analysis")

	flag.Usage = usage
	flag.Parse()

	commandContext.ImportSpec = flag.Arg(0)

	log.Printf("Command context: %s\n", os.Args)

	return commandContext
}
