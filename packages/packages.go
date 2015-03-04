package packages

import (
	"go/parser"
	"go/token"
	"go/ast"
	"log"
)

func Packages(files []string) map[string]*ast.Package {
	fset := token.NewFileSet()

	allPackages, err := ParseFiles(fset, files, 0)
	if err != nil {
		log.Fatal(err)
	}

	for name, pkg := range allPackages {
		log.Printf("\nPackage: %#v", pkg)
		log.Printf("\n*** Package***\nName: %s\nFiles: %#v\n\n", name, pkg.Files)
	}

	return allPackages
}

func ParseFiles(fset *token.FileSet, files []string, mode parser.Mode) (pkgs map[string]*ast.Package, first error) {
	pkgs = make(map[string]*ast.Package)
	
	for _, filename := range files {
		if src, err := parser.ParseFile(fset, filename, nil, mode); err == nil {
			name := src.Name.Name
			pkg, found := pkgs[name]
			if !found {
				pkg = &ast.Package{
					Name:  name,
					Files: make(map[string]*ast.File),
				}
				pkgs[name] = pkg
			}
			pkg.Files[filename] = src
		} else if first == nil {
			first = err
		}
	}

	return
}