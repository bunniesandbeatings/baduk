package parser
import (
	"go/build"
	. "github.com/bunniesandbeatings/go-flavor-parser/ast"
	"strings"
	"go/token"
	"go/parser"
	"log"
	"go/ast"
	"github.com/bunniesandbeatings/go-flavor-parser/parser/first"
	"path/filepath"
)

func UpdateFromPackage(pkg *build.Package, root *Directory) {
	fset := token.NewFileSet()

	dir := getOrCreateDirectory(pkg.ImportPath, root)

//	receiverFunctions = make()

	for _, filename := range pkg.GoFiles {
		filepath := filepath.Join(pkg.Dir + "/" + filename)
		astFile, err := parser.ParseFile(fset, filepath, nil, 0)

//		spew.Dump(astFile)

		if err != nil {
			log.Printf("Error %s when parsing file %s\n", err, filepath)
			dir.Files[filename] = nil
		} else {
			dir.Files[filename] = &File{}

			rootVisitor := first.RootVisitor{File: dir.Files[filename] }
			ast.Walk(rootVisitor, astFile)
		}
	}
}

func getOrCreateDirectory(path string, root *Directory) *Directory {
	pathSections := strings.Split(path, "/")

	currentNode := root

	for _, pathSection := range pathSections {
		if _, found := currentNode.Directories[pathSection]; !found {
			currentNode.Directories[pathSection] = NewDirectory()
		}

		currentNode = currentNode.Directories[pathSection]
	}

	return currentNode
}