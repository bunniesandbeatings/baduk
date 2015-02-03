package parser

import (
	. "github.com/bunniesandbeatings/go-flavor-parser/packages"
	. "github.com/bunniesandbeatings/go-flavor-parser/structure101/datafiles"
	"encoding/xml"

	"fmt"
)

type Parser struct {
	dataFile DataFile
	packages Packages
}

func NewParser() *Parser {
	parser := Parser{}
	
	parser.dataFile = *NewDataFile("com.bunniesandbeatings.go-flavor")

	parser.packages = *NewPackages()
	// TODO make a NewPackageHandler interface
	parser.packages.NewPackageHandler = parser.createNewPackageHandler()

	return &parser
}

func (parser *Parser) AddImports(importPaths []string) {
	parser.packages.AddByImportPaths(importPaths, false)
}

func (parser *Parser) createNewPackageHandler() func(Package) {
	return func(packageDef Package) {
		newModule := Module{
			Name: packageDef.ImportPath,
			Id:   packageDef.UniqueId(),
			Type: "package",
		}
		
		parser.dataFile.Modules = append(parser.dataFile.Modules, newModule)
	}
}

func (parser *Parser) DataFileXML() []byte {
	dataFileXML, err := xml.MarshalIndent(parser.dataFile, "", "  ")
	if err != nil {
		fmt.Printf("error when Marshalling Data File definition: %v\n", err)
	}
	return dataFileXML
}