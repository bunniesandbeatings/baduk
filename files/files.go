package files

import (
	"github.com/mndrix/ps"

	"fmt"
	"go/build"
)

func GetFiles(buildContext build.Context, visitedPackages ps.Map, importPaths []string) (newVisitedPackages ps.Map, files []string) {
	if visitedPackages == nil {
		visitedPackages = ps.NewMap()
	}

	for _, importPath := range importPaths {
		if _, found := visitedPackages.Lookup(importPath); !found {
			buildPackage, err := buildContext.Import(importPath, ".", 0)

			visitedPackages = visitedPackages.Set(importPath, new(ps.Any))

			if err == nil {
				files = mergeFiles(files, buildPackage.Dir, buildPackage.GoFiles)

				newVisitedPackages, filesFromImports := GetFiles(buildContext, visitedPackages, buildPackage.Imports)

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
