package files

import (
	"github.com/mndrix/ps"
	"github.com/bunniesandbeatings/gotool"

	"fmt"
	"go/build"
)

func GetFilesFromImportSpec(buildContext build.Context, importSpec string) []string {
	gotool.SetContext(buildContext)
	importPaths := gotool.ImportPaths([]string{importSpec})

	visitedPackages := ps.NewMap()
	_, files := getFilesFromImportPaths(buildContext, visitedPackages, importPaths)
	
	return files
}


func getFilesFromImportPaths(buildContext build.Context, visitedPackages ps.Map, importPaths []string) (newVisitedPackages ps.Map, files []string) {
	for _, importPath := range importPaths {
		if _, found := visitedPackages.Lookup(importPath); !found {
			buildPackage, err := buildContext.Import(importPath, ".", 0)

			visitedPackages = visitedPackages.Set(importPath, new(ps.Any))

			if err == nil {
				files = mergeFiles(files, buildPackage.Dir, buildPackage.GoFiles)

				newVisitedPackages, filesFromImports := getFilesFromImportPaths(buildContext, visitedPackages, buildPackage.Imports)

				visitedPackages = newVisitedPackages

				files = append(files, filesFromImports...)
			} else {
				fmt.Printf("WARNING: could not get file list for '%s', go/build#import failed with: %s", importPath, err)
			}
		}
	}

	return visitedPackages, files
}

func mergeFiles(filesWithPaths []string, importPath string, files []string) (newFilesWithPaths []string) {
	for _, filename := range files {
		filesWithPaths = append(filesWithPaths, importPath+"/"+filename)
	}

	return filesWithPaths
}
