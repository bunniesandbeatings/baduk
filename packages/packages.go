package packages

import (
	"fmt"
	"go/build"
)

type Package struct {
	*build.Package
}

func CreatePackage(packageName string) Package {
	buildPackage, err := build.Import(packageName, ".", 0)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	packageDef := Package{
		buildPackage,
	}

	return packageDef
}
