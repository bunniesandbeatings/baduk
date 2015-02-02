package packages

import (
	"fmt"
	"go/build"
)

type Package struct {
	*build.Package
}

func CreatePackage(importPath string) Package {
	buildPackage, err := build.Import(importPath, ".", 0)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	packageDef := Package{
		buildPackage,
	}

	return packageDef
}
