package main

import (
	"flag"
	"fmt"
	"github.com/kisielk/gotool"
	"os"

	. "github.com/bunniesandbeatings/go-flavor-parser/packages"
)

func usage() {
	appName := os.Args[0]
	fmt.Fprintf(os.Stderr, "%s usage:\n", appName)
	fmt.Fprintf(os.Stderr, "\t%s [flags] packages # see 'go help packages'\n", appName)
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	importSpec := []string{flag.Arg(0)}

	packages := []Package{}

	for _, packageName := range gotool.ImportPaths(importSpec) {
		fmt.Printf("Parsing: %s\n", packageName)
		packageDef := CreatePackage(packageName)
		packages = append(packages, packageDef)
	}

	for _, def := range packages {
		fmt.Printf("Package: %s\n", def.Name)
	}
}
