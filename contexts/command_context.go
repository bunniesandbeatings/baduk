package contexts

import (
	"github.com/codegangsta/cli"
	"os"
)

type CommandContext struct {
	GoRoot     string
	GoPath     string
	ImportSpec string
}

func CreateCommandContextFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "gopath",
			Value: os.Getenv("GOPATH"),
			Usage: "allows you to choose a different GOPATH to use during analysis",
		},
		cli.StringFlag{
			Name:  "goroot",
			Value: os.Getenv("GOROOT"),
			Usage: "allows you to choose a different GOROOT to use during analysis",
		},
	}
}

func CreateCommandContext(cliContext *cli.Context) CommandContext {
	commandContext := CommandContext{}

	commandContext.GoRoot = cliContext.String("goroot")
	commandContext.GoPath = cliContext.String("gopath")
	commandContext.ImportSpec = cliContext.Args()[0]

	return commandContext
}
