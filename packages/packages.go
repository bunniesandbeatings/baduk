package packages

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

	if !found {
		packageDef := CreatePackage(importPath)
	
		packages.byImportPath[importPath] = *packageDef

		if packages.NewPackageHandler != nil {
			packages.NewPackageHandler(*packageDef)
		}
		
		if recur{
			packages.AddByImportPaths(packageDef.Imports, true)
		}
	}

}

