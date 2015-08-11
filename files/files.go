package files

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"

	"github.com/bunniesandbeatings/gotool"
	"github.com/mndrix/ps"
	"github.com/davecgh/go-spew/spew"
)

func Files(buildContext build.Context, importSpec string) ps.Map {
	importPaths := importPaths(buildContext, importSpec)

	fmt.Println(">>>> DEBUG: ImportPaths")
	spew.Dump(importPaths)
	fmt.Println("<<<< DEBUG: ImportPaths\n\n")

	files := ps.NewMap()
	files = filesFromImportPaths(buildContext, files, importPaths)

	return files
}

func importPaths(buildContext build.Context, importSpec string) []string {
	gotool.SetContext(buildContext)
	return gotool.ImportPaths([]string{importSpec})
}

func filesFromImportPaths(buildContext build.Context, filesByImportPath ps.Map, importPaths []string) (newFilesByImportPath ps.Map) {
	fset := token.NewFileSet()

	newFilesByImportPath = filesByImportPath

	for _, importPath := range importPaths {

		if _, found := newFilesByImportPath.Lookup(importPath); !found {
			fmt.Println("DEBUG: Handling path: %s", importPath)
			buildPackage, err := buildContext.Import(importPath, ".", 0)

			if err != nil {
				log.Print(fmt.Sprintf("WARNING: Could not get file list for '%s', import failed with: %s", importPath, err))
				newFilesByImportPath = newFilesByImportPath.Set(importPath, []*ast.File{})
			} else {
				files := mergeFiles(buildPackage.Dir, buildPackage.GoFiles)

				asts := []*ast.File{}

				for _, filename := range files {
					fileAst, err := parser.ParseFile(fset, filename, nil, parser.ImportsOnly)

					if err != nil {
						log.Fatal(err)
					}

					asts = append(asts, fileAst)
				}

				newFilesByImportPath = newFilesByImportPath.Set(importPath, asts)

				newFilesByImportPath = filesFromImportPaths(buildContext, newFilesByImportPath, buildPackage.Imports)
			}
		}
	}

	return
}

func mergeFiles(importPath string, files []string) (filesWithPaths []string) {
	filesWithPaths = []string{}
	for _, filename := range files {
		filesWithPaths = append(filesWithPaths, importPath+"/"+filename)
	}

	return
}
