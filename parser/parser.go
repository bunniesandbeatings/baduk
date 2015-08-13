package parser

import (
	. "github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"github.com/bunniesandbeatings/go-flavor-parser/parser/first"
	"github.com/bunniesandbeatings/gotool"
	"go/ast"
	"go/build"
	goparser "go/parser"
	"go/token"
	"log"
	"path/filepath"
)

type Parser struct {
	Arch    *Architecture
	Context build.Context
}

func NewParser(context build.Context) *Parser {
	return &Parser{
		Arch:    NewArchitecture(),
		Context: context,
	}
}

func (parser *Parser) ParseImportSpec(spec string) {
	gotool.SetContext(parser.Context)
	importPaths := gotool.ImportPaths([]string{spec})

	for _, importPath := range importPaths {
		parser.ParseImport(importPath)
	}
}

func (parser *Parser) ParseImport(importPath string) {
	log.Printf("Parsing import path: %s", importPath)

	buildPackage, _ := parser.Context.Import(importPath, ".", 0)

	parser.updateFromPackage(buildPackage)
}

func (parser *Parser) updateFromPackage(pkg *build.Package) {
	fset := token.NewFileSet()

	dir := parser.Arch.FindDirectory(pkg.ImportPath)

	//	receiverFunctions = make()

	for _, filename := range pkg.GoFiles {
		filepath := filepath.Join(pkg.Dir + "/" + filename)
		astFile, err := goparser.ParseFile(fset, filepath, nil, 0)

		//		spew.Dump(astFile)

		if err != nil {
			log.Printf("Error %s when parsing file %s\n", err, filepath)
			dir.Files[filename] = nil
		} else {
			dir.Files[filename] = &File{}

			rootVisitor := first.RootVisitor{File: dir.Files[filename]}
			ast.Walk(rootVisitor, astFile)
		}
	}
}
