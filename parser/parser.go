package parser
import (
	"go/build"
	target "github.com/bunniesandbeatings/go-flavor-parser/ast"
	"strings"
	"go/token"
	"go/parser"
	"github.com/davecgh/go-spew/spew"
	"log"
	"go/ast"
	"github.com/bunniesandbeatings/go-flavor-parser/parser/first"
)

func UpdateFromPackage(pkg *build.Package, target_ast *target.PathNode) {

	path := strings.Split(pkg.ImportPath, "/")

	currentNode := target_ast

	for _, pathSection := range path {
		if _, found := currentNode.PathChildren[pathSection]; !found {
			currentNode.PathChildren[pathSection] = target.NewPathNode()
		}

		currentNode = currentNode.PathChildren[pathSection]
	}

	fset := token.NewFileSet()

	currentNode.FileChildren = make(map[string]*target.FileNode)

	for _, filename := range pkg.GoFiles {
		// TODO: whats the portable way?
		filepath := pkg.Dir + "/" + filename
		astFile, err := parser.ParseFile(fset, filepath, nil, 0)

		spew.Dump(astFile)

		if err != nil {
			log.Printf("Error %s when parsing file %s\n", err, filepath)
			currentNode.FileChildren[filename] = nil
		} else {
			currentNode.FileChildren[filename] = &target.FileNode{}

			rootVisitor := first.RootVisitor{File: currentNode.FileChildren[filename] }
			ast.Walk(rootVisitor, astFile)
		}

	}
}
