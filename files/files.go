package files

import (
	"github.com/mndrix/ps"
	"github.com/bunniesandbeatings/gotool"
	
	"go/ast"
	"go/token"
	"go/parser"
	"log"
	"go/build"
)

func ImportPaths(buildContext build.Context, importSpec string) []string {
	gotool.SetContext(buildContext)
	return gotool.ImportPaths([]string{importSpec})
}

func Files(buildContext build.Context, importSpec string) ps.Map {
	importPaths := ImportPaths(buildContext, importSpec)

	files := ps.NewMap()
	files = FilesFromImportPaths(buildContext, files, importPaths)

	return files
}

func FilesFromImportPaths(buildContext build.Context, filesByImportPath ps.Map, importPaths []string) (newFilesByImportPath ps.Map) {
	fset := token.NewFileSet()
	
	newFilesByImportPath = filesByImportPath

	for _, importPath := range importPaths {
		if _, found := filesByImportPath.Lookup(importPath); !found {
			buildPackage, err := buildContext.Import(importPath, ".", 0)

			if err != nil {
				log.Fatal("Could not get file list for '%s', go/build#import failed with: %s", importPath, err)
			}

			files := MergeFiles(buildPackage.Dir, buildPackage.GoFiles)
			
			asts := []*ast.File{}
			
			for _, filename := range files {
				fileAst, err := parser.ParseFile(fset, filename, nil, parser.ImportsOnly)
				
				if err != nil {
					log.Fatal(err)
				}

				asts = append(asts, fileAst)
			}

			newFilesByImportPath = newFilesByImportPath.Set(importPath, asts)

			newFilesByImportPath = FilesFromImportPaths(buildContext, newFilesByImportPath, buildPackage.Imports)
		}
	}

	return
}

func MergeFiles(importPath string, files []string) (filesWithPaths []string) {
	filesWithPaths = []string{}
	for _, filename := range files {
		filesWithPaths = append(filesWithPaths, importPath+"/"+filename)
	}

	return
}
