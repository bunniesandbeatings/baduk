package packages

import (
	"crypto/md5"
	"encoding/hex"
	"go/build"
	"io"
)

type Package struct {
	*build.Package
}

func CreatePackage(importPath string, buildContext build.Context) (*Package, error) {
	buildPackage, err := buildContext.Import(importPath, ".", 0)

	if err != nil {
		return nil, err
	}

	packageDef := Package{
		buildPackage,
	}

	return &packageDef, err
}

func (packageDef *Package) UniqueId() string {
	hash := md5.New()
	io.WriteString(hash, packageDef.ImportPath)

	return hex.EncodeToString(hash.Sum([]byte{}))
}
