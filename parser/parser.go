package parser

import (
	"fmt"
	"go/ast"
	"go/build"
	goparser "go/parser"
	"go/token"
	"log"
	"path/filepath"

	arch "github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"github.com/bunniesandbeatings/go-flavor-parser/parser/first"
	"github.com/bunniesandbeatings/gotool"
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
	buildPackage, _ := parser.context.Import(importPath, ".", 0)
	parser.parsePackage(buildPackage)
}

func (parser *Parser) parsePackage(pkg *build.Package) {
	dirNode := parser.arch.FindDirectory(pkg.ImportPath)

	astFiles := make([]*ast.File, 1)
	for _, filename := range pkg.GoFiles {
		filepath := filepath.Join(pkg.Dir, filename)
		astFile := parser.getASTFile(filepath)
		astFiles = append(astFiles, astFile)
	}

	for _, filename := range pkg.GoFiles {
		parser.parseGoFile(dirNode, pkg.Dir, filename, astFiles)
	}
}

func (parser *Parser) parseGoFile(directory *arch.Directory, path string, filename string, astFiles []*ast.File) {
	filepath := filepath.Join(path, filename)
	astFile := parser.getASTFile(filepath)
	//	spew.Dump(astFile)

	pkg, err := directory.CreatePackage(astFile.Name.Name, astFiles)
	if err != nil {
		panic(fmt.Sprintf("File %s package %s conflicts with %s", filepath, astFile.Name.Name, directory.Package))
	}

	visitorContext := first.Context{
		Package:  pkg,
		Filename: filename,
		Fset:     parser.fset,
	}

	rootVisitor := first.NewRootVisitor(visitorContext)

	ast.Walk(rootVisitor, astFile)
}

func (parser *Parser) getASTFile(filepath string) (astFile *ast.File) {
	astFile, err := goparser.ParseFile(parser.fset, filepath, nil, 0)
	if err != nil {
		log.Printf("WARNING: Error %s when parsing file %s\n", err, filepath)
		astFile = nil
	}
	return
}
