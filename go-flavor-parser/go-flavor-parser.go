package main

import (
	"github.com/bunniesandbeatings/go-flavor-parser/files"
	"github.com/bunniesandbeatings/go-flavor-parser/context"

	"flag"
	"log"
	"os"
	"go/token"
	. "go/parser"
	"go/ast"
	"strings"
)

func usage() {
	appName := os.Args[0]
	log.Printf("%s usage:\n", appName)
	log.Printf("\t%s [flags] packages # see 'go help packages'\n", appName)
	log.Printf("Flags:\n")
	flag.PrintDefaults()
}

func main() {
	commandContext := context.CreateCommandContext(usage)
	buildContext := context.CreateBuildContext(commandContext)

	allFiles := files.GetFilesFromImportSpec(buildContext, commandContext.ImportSpec)

	log.Printf("\n*** Files to process ***\n  %#s\n\n", strings.Join(allFiles, "\n  "))
	
	fset := token.NewFileSet()

	allPackages, err := ParseFiles(fset, allFiles, 0)
	if err != nil {
		log.Println(err)
		return
	}

	for name, pkg := range allPackages {
		log.Printf("\n*** Package***\nName: %s\nFiles: %#v\n\n", name, pkg.Files)
	}
}

func ParseFiles(fset *token.FileSet, files []string, mode Mode) (pkgs map[string]*ast.Package, first error) {
	pkgs = make(map[string]*ast.Package)
	for _, filename := range files {
		if src, err := ParseFile(fset, filename, nil, mode); err == nil {
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
