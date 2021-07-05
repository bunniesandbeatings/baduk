package cmd

import (
	"errors"
	"go/build"
	"log"
	"os"

	badukParser "github.com/bunniesandbeatings/baduk/parser"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse [packageSpec]",
	Short: "parse a package",
	Long:  "parse a package",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a packageSpec argument")
		}
		// FIXME validate package spec here?
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		buildContext := build.Default

		buildContext.GOPATH = goPath
		buildContext.GOROOT = os.Getenv("GOROOT")

		parser := badukParser.NewParser(buildContext)

		parser.ParseImportSpec(args[0])

		log.Println(spew.Sdump(parser.GetArchitecture()))
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
