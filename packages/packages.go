package packages

type Packages struct {
	ByImportPath map[string]Package
}

func CreatePackages(importPaths []string) Packages {
	packages := Packages{
		ByImportPath: make(map[string]Package),
	}
	packages.AddPackagesByImportPaths(importPaths)

	return packages
}

func (packages *Packages) AddPackageByImportPath(importPath string) Package {
	var packageDef Package

	if packageDef, present := packages.ByImportPath[importPath]; !present {
		packageDef = CreatePackage(importPath)
		packages.ByImportPath[importPath] = packageDef
	}

	return packageDef
}

func (packages *Packages) AddPackagesByImportPaths(importPaths []string) {
	for _, importPath := range importPaths {
		packages.AddPackageByImportPath(importPath)
	}
}

func (packages *Packages) ExpandImportedPackages() {
	// Avoid expanding the map while iterating
	currentPackageList := []string{}

	for importPath, _ := range packages.ByImportPath {
		currentPackageList = append(currentPackageList, importPath)
	}

	for _, importPath := range currentPackageList {
		for _, dependencyImportPath := range packages.ByImportPath[importPath].Imports {
			packages.AddPackageByImportPath(dependencyImportPath)
		}
	}

}
