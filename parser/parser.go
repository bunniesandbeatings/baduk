package parser

import (
	. "github.com/bunniesandbeatings/go-flavor-parser/packages"
	. "github.com/bunniesandbeatings/go-flavor-parser/structure101/datafiles"
	"encoding/xml"

	"fmt"
)

type Parser struct {
	Datafile DataFile
	Packages Packages
}

func NewParser() *Parser {
	parser := Parser{}

	parser.Datafile = *NewDataFile("com.bunniesandbeatings.go-flavor")

	parser.Packages = *NewPackages()
	// TODO make a NewPackageHandler interface
	parser.Packages.NewPackageHandler = parser.createNewPackageHandler()

	return &parser
}

func (parser *Parser) AddImportPaths(importPaths []string) {
	parser.Packages.AddByImportPaths(importPaths, true)
}

func (parser *Parser) createNewPackageHandler() func(Package) {
	return func(packageDef Package) {
		newModule := Module{
			Name: packageDef.ImportPath,
			Id:   packageDef.UniqueId(),
			Type: "package",
		}

		parser.Datafile.Modules = append(parser.Datafile.Modules, newModule)

		for _, dependencyImportPath := range packageDef.Imports {
			importPackage, found := parser.Packages.FindByImportPath(dependencyImportPath)

			if found {
				newDependency := Dependency{
					From: packageDef.UniqueId(),
					To: importPackage.UniqueId(),
					Type: "imports",
				}
				parser.Datafile.Dependencies = append(parser.Datafile.Dependencies, newDependency)
			}
		}
	}
}

func (parser *Parser) DataFileXML() []byte {
	dataFileXML, err := xml.MarshalIndent(parser.Datafile, "", "  ")
	if err != nil {
		fmt.Printf("error when Marshalling Data File definition: %v\n", err)
	}
	return dataFileXML
}