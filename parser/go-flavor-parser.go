package main

import (
	"flag"
	"fmt"
	"github.com/kisielk/gotool"
	"os"

	"encoding/xml"
	. "github.com/bunniesandbeatings/go-flavor-parser/packages"
	. "github.com/bunniesandbeatings/go-flavor-parser/structure101/datafiles"
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
	projectImportPaths := gotool.ImportPaths(importSpec)

	packages := CreatePackages(projectImportPaths)

	packages.ExpandImportedPackages()

	dataFile := packagesToDataFile(packages)

	output, err := xml.MarshalIndent(dataFile, "", "  ")
	if err != nil {
		fmt.Printf("error when Marshalling Data File definition: %v\n", err)
	}

	fmt.Println(string(output))
}

func packagesToDataFile(packages Packages) *DataFile {
	modules := []Module{}

	for name, packageDef := range packages.ByImportPath {
		newModule := Module{
			Name: name,
			Id:   packageDef.UniqueID(),
			Type: "package",
		}

		modules = append(modules, newModule)
	}

	return &DataFile{
		Flavor:  "com.bunniesandbeatings.go-flavor",
		Modules: modules,
	}
}
