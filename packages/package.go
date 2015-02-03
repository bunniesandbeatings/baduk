package packages

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go/build"
	"io"
)

type Package struct {
	*build.Package
}

func CreatePackage(importPath string) *Package {
	buildPackage, err := build.Import(importPath, ".", 0)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	packageDef := Package{
		buildPackage,
	}

	return &packageDef
}

func (packageDef *Package) UniqueId() string {
	hash := md5.New()
	io.WriteString(hash, packageDef.ImportPath)

	return hex.EncodeToString(hash.Sum([]byte{}))
}
