package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"

	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"runtime/debug"

	"github.com/bunniesandbeatings/gotool"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func NewPathNode() (pathNode *PathNode) {
	pathNode = &PathNode{
		Children: make(map[string]*PathNode),
	}

	return
}

type PathNode struct {
	Children map[string]*PathNode
	Package  string
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

func mergeFiles(importPath string, files []string) (filesWithPaths []string) {
	filesWithPaths = []string{}
	for _, filename := range files {
		filesWithPaths = append(filesWithPaths, importPath+"/"+filename)
	}

	return
}

func updateASTWithPackage(pkg *build.Package, ast *PathNode) {

	path := strings.Split(pkg.ImportPath, "/")

	currentNode := ast

	for _, pathSection := range path {
		if _, found := currentNode.Children[pathSection]; !found {
			currentNode.Children[pathSection] = NewPathNode()
		}

		currentNode = currentNode.Children[pathSection]
	}

	currentNode.Package = pkg.Name
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

//	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> DEBUG: ImportPaths")
//	spew.Dump(importPaths)
//	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< DEBUG: ImportPaths\n\n")

	//	fset := token.NewFileSet()

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
