package parser

import (
	arch "github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"github.com/bunniesandbeatings/gotool"
	"go/build"
	"go/token"
	"log"
	"path/filepath"
	goparser "go/parser"
	"github.com/bunniesandbeatings/go-flavor-parser/parser/first"
	"go/ast"
)

type Parser struct {
	Arch    *arch.Architecture
	Context build.Context
	fset    *token.FileSet
}

func NewParser(context build.Context) *Parser {
	return &Parser{
		Arch:    arch.NewArchitecture(),
		Context: context,
		fset: token.NewFileSet(),
	}
}

func (parser *Parser) ParseImportSpec(spec string) {
	gotool.SetContext(parser.Context)
	importPaths := gotool.ImportPaths([]string{spec})

	for _, importPath := range importPaths {
		parser.parseImport(importPath)
	}
}

func (parser *Parser) parseImport(importPath string) {
	log.Printf("Parsing import path: %s", importPath)

	buildPackage, _ := parser.Context.Import(importPath, ".", 0)

	parser.parsePackage(buildPackage)
}

func (parser *Parser) parsePackage(pkg *build.Package) {
	dirNode := parser.Arch.FindDirectory(pkg.ImportPath)

	for _, filename := range pkg.GoFiles {
		parser.parseGoFile(dirNode, pkg.Dir, filename)
	}
}

func (parser *Parser) parseGoFile(dir *arch.Directory, path string, filename string) {
	filepath := filepath.Join(path, filename)

	astFile, err := goparser.ParseFile(parser.fset, filepath, nil, 0)
	if err != nil {
		log.Printf("WARNING: Error %s when parsing file %s\n", err, filepath)
		dir.Files[filename] = nil
		return
	}

	dir.Files[filename] = &arch.File{}
	//		spew.Dump(astFile)

	rootVisitor := first.NewRootVisitor(dir.Files[filename])
	ast.Walk(rootVisitor, astFile)
}
