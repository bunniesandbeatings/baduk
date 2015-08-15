package parser

import (
	"fmt"
	arch "github.com/bunniesandbeatings/go-flavor-parser/architecture"
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
	arch    *arch.Architecture
	context build.Context
	fset    *token.FileSet
}

func NewParser(context build.Context) *Parser {
	return &Parser{
		arch:    arch.NewArchitecture(),
		context: context,
		fset:    token.NewFileSet(),
	}
}

func (parser Parser) GetArchitecture() arch.Architecture {
	return *parser.arch
}

func (parser *Parser) ParseImportSpec(spec string) {
	gotool.SetContext(parser.context)
	importPaths := gotool.ImportPaths([]string{spec})

	for _, importPath := range importPaths {
		parser.parseImport(importPath)
	}
}

func (parser *Parser) parseImport(importPath string) {
	log.Printf("Parsing import path: %s", importPath)

	buildPackage, _ := parser.context.Import(importPath, ".", 0)

	parser.parsePackage(buildPackage)
}

func (parser *Parser) parsePackage(pkg *build.Package) {
	dirNode := parser.arch.FindDirectory(pkg.ImportPath)

	for _, filename := range pkg.GoFiles {
		parser.parseGoFile(dirNode, pkg.Dir, filename)
	}
}

func (parser *Parser) parseGoFile(dir *arch.Directory, path string, filename string) {
	filepath := filepath.Join(path, filename)

//	packageMaps := PackageMaps{}

	astFile, err := goparser.ParseFile(parser.fset, filepath, nil, 0)
	if err != nil {
		log.Printf("WARNING: Error %s when parsing file %s\n", err, filepath)
		dir.Files[filename] = nil
		return
	}

	packageName := astFile.Name.Name

	if dir.Package == "" {
		dir.Package = packageName
	} else if dir.Package != packageName {
		panic(fmt.Sprintf("File %s package %s conflicts with %s", filepath, packageName, dir.Package))
	}

	file := &arch.File{}

	dir.Files[filename] = file

	//	spew.Dump(astFile)

	rootVisitor := first.RootVisitor{File: file}
	ast.Walk(rootVisitor, astFile)
}
