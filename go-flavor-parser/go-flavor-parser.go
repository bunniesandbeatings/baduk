package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"

	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"runtime/debug"

	"strings"

	"go/parser"
	"go/token"

	"github.com/bunniesandbeatings/gotool"
	"github.com/davecgh/go-spew/spew"
	. "go/ast"
	"reflect"
)

type AFakeInterface interface {
	Foo() string
}

type BFakeInterface interface {
	AFakeInterface
	Bar() string
}

type FileNode struct {
	PublicStructs    []string
	PublicInterfaces []string
	PublicFuncs      []string
}


func NewPathNode() (pathNode *PathNode) {
	pathNode = &PathNode{
		PathChildren: make(map[string]*PathNode),
	}

	return
}


type PathNode struct {
	PathChildren map[string]*PathNode
	FileChildren map[string]*FileNode
}

// VISITORS

type DumpVisitor struct {
}

func (visitor DumpVisitor) Visit(node Node) Visitor {
	fmt.Println(reflect.TypeOf(node))
	return visitor
}


type TypeSpecVisitor struct {
	File     *FileNode
	TypeSpec *TypeSpec
}

func (visitor TypeSpecVisitor) Visit(node Node) Visitor {
	// TODO: Can types be private?

	switch node.(type) {
	case *InterfaceType:
		visitor.File.PublicInterfaces = append(visitor.File.PublicInterfaces, visitor.TypeSpec.Name.Name)
	case *StructType:
		visitor.File.PublicStructs = append(visitor.File.PublicStructs, visitor.TypeSpec.Name.Name)
	}
	return nil
}

type GenDeclVisitor struct {
	File *FileNode
}

func (visitor GenDeclVisitor) Visit(node Node) Visitor {
	switch t := node.(type) {
	case *TypeSpec:
		return TypeSpecVisitor{
			File: visitor.File,
			TypeSpec: t,
		}
	}
	return nil
}

type RootVisitor struct {
	File *FileNode
}

func (visitor RootVisitor) Visit(node Node) Visitor {
	switch t := node.(type) {
	case *GenDecl:
		return GenDeclVisitor{
			File: visitor.File,
		}
	case *FuncDecl:
		// TODO: filter public only
		if t.Recv == nil {
			visitor.File.PublicFuncs = append(visitor.File.PublicFuncs, t.Name.Name)
		} else {
			// queue function with receiver
		}
	}
	return visitor
}

func usage() {
	appName := os.Args[0]
	log.Printf("%s usage:\n", appName)
	log.Printf("\t%s [flags] packages # see 'go help packages'\n", appName)
	log.Printf("Flags:\n")
	flag.PrintDefaults()
}

func importPaths(buildContext build.Context, importSpec string) []string {
	gotool.SetContext(buildContext)
	return gotool.ImportPaths([]string{importSpec})
}

func updateASTWithPackage(pkg *build.Package, ast *PathNode) {

	path := strings.Split(pkg.ImportPath, "/")

	currentNode := ast

	for _, pathSection := range path {
		if _, found := currentNode.PathChildren[pathSection]; !found {
			currentNode.PathChildren[pathSection] = NewPathNode()
		}

		currentNode = currentNode.PathChildren[pathSection]
	}

	fset := token.NewFileSet()

	currentNode.FileChildren = make(map[string]*FileNode)

	for _, filename := range pkg.GoFiles {
		// TODO: whats the portable way?
		filepath := pkg.Dir + "/" + filename
		astFile, err := parser.ParseFile(fset, filepath, nil, 0)

		spew.Dump(astFile)

		if err != nil {
			log.Printf("Error %s when parsing file %s\n", err, filepath)
			currentNode.FileChildren[filename] = nil
		} else {
			currentNode.FileChildren[filename] = &FileNode{}

			Walk(RootVisitor{File: currentNode.FileChildren[filename] }, astFile)

		}

	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			debug.PrintStack()
			os.Exit(1)
		}
	}()

	commandContext := contexts.CreateCommandContext(usage)

	buildContext := contexts.CreateBuildContext(commandContext)

	ast := NewPathNode()

	importPaths := importPaths(buildContext, commandContext.ImportSpec)


	for _, importPath := range importPaths {
		fmt.Println(importPath)

		buildPackage, _ := buildContext.Import(importPath, ".", 0)

		//		spew.Dump(buildPackage)

		updateASTWithPackage(buildPackage, ast)
	}

	fmt.Println(">>>> DEBUG: ast")
	spew.Dump(ast)
	fmt.Println("<<<< DEBUG: ast\n\n")

	//	datafile := datafiles.NewDataFile("com.bunniesandbeatings.go-flavor")
	//
	//	allFiles.ForEach(func(importPath string, files Any) {
	//		module := datafiles.Module{
	//			Name: importPath,
	//			Id: importPath,
	//			Type: "package",
	//		}
	//
	//		datafile.Modules = append(datafile.Modules, module)
	//		log.Printf("%s:", importPath)
	//
	//		for _, file := range files.([]*ast.File) {
	//			log.Printf("\t%#v:", *file)
	//		}
	//
	//	})

	//	ioutil.WriteFile(commandContext.OutputPath, datafile.ToXML(), 0644)
	//
	//	fmt.Printf("Output written to '%s'\n", commandContext.OutputPath)
	//	_ = packages.Packages(allFiles)
}
