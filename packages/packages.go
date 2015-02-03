package packages

import "fmt"

type Packages struct {
	byImportPath map[string]Package
	NewPackageHandler func(Package)
}

func NewPackages() *Packages {
	return &Packages{
		byImportPath: make(map[string]Package),
	}
}

func (packages *Packages) AddByImportPaths(importPaths []string, recur bool) {
	for _, importPath := range importPaths {
		packages.AddByImportPath(importPath, recur)
	}
}

func (packages *Packages) AddByImportPath(importPath string, recur bool) {
	_, found := packages.byImportPath[importPath]

	if found {
		return
	}

	packageDef, err := CreatePackage(importPath)

	if err != nil {
		fmt.Printf("Could not find import (skipping): '%s'\n", importPath)
		packages.byImportPath[importPath] = Package{}
		return
	}

	packages.byImportPath[importPath] = *packageDef
	packages.handleNewPackage(*packageDef)
	if recur {
		packages.AddByImportPaths(packageDef.Imports, true)
	}
}

func (packages *Packages) handleNewPackage(packageDef Package) {
	if packages.NewPackageHandler == nil {
		return
	}

	packages.NewPackageHandler(packageDef)
}