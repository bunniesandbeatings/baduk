package parser
import (
	"go/build"
	. "github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"go/token"
	"go/parser"
	"log"
	"go/ast"
	"github.com/bunniesandbeatings/go-flavor-parser/parser/first"
	"path/filepath"
)

func UpdateFromPackage(pkg *build.Package, arch *Architecture) {
	fset := token.NewFileSet()

	dir := arch.FindDirectory(pkg.ImportPath)

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

