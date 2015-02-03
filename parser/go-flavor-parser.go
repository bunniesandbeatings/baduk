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
	projectImportPaths := importPathsFromCommandLine()
	
	dataFile := NewDataFile("com.bunniesandbeatings.go-flavor")

	packages := NewPackages()
	packages.NewPackageHandler = CreateNewPackageHandler(&dataFile)

	packages.AddByImportPaths(projectImportPaths, false)

	dataFileXML := marshalDatafile(dataFile)

	fmt.Println(string(dataFileXML))
}

func CreateNewPackageHandler(datafile *DataFile) func(Package) {
	return func(packageDef Package) {
		newModule := Module{
			Name:
			packageDef.ImportPath,
			Id:   packageDef.UniqueId(),
			Type: "package",
		}

		datafile.Modules = append(datafile.Modules, newModule)
	}
}

func importPathsFromCommandLine() []string {
	flag.Usage = usage
	flag.Parse()

	importSpec := []string{flag.Arg(0)}

	return gotool.ImportPaths(importSpec)
}

func marshalDatafile(dataFile DataFile) []byte {
	dataFileXML, err := xml.MarshalIndent(dataFile, "", "  ")
	if err != nil {
		fmt.Printf("error when Marshalling Data File definition: %v\n", err)
	}
	return dataFileXML
}