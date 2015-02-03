package packages

type Packages struct {
	ByImportPath map[string]Package
	NewPackageHandler func(Package)
}

func NewPackages() Packages {
	return Packages{
		ByImportPath: make(map[string]Package),
	}
}

func (packages *Packages) AddByImportPaths(importPaths []string, recur bool) {
	for _, importPath := range importPaths {
		packages.AddByImportPath(importPath, recur)
	}

}

func (packages *Packages) AddByImportPath(importPath string, recur bool) (packageDef Package, found bool) {
	packageDef, found = packages.ByImportPath[importPath]

	if !found {
		packageDef = CreatePackage(importPath)
		
		if packages.NewPackageHandler != nil {
			packages.NewPackageHandler(packageDef)
		}
		
		if recur{
			packages.AddByImportPaths(packageDef.Imports, true)
		}
	}

	return packageDef, found
}

