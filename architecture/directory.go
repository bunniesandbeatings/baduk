package architecture

import (
	"errors"
	"go/ast"
)

type Directory struct {
	Directories map[string]*Directory
	Package     *Package
}

func NewDirectory() *Directory {
	return &Directory{
		Directories: make(map[string]*Directory),
	}
}

func (directory *Directory) CreatePackage(name string, astFiles []*ast.File) (pkg *Package, err error) {
	if directory.Package == nil {
		pkg = NewPackage(name, astFiles)
		directory.Package = pkg
	} else if directory.Package.Name != name {
		err = errors.New("Cannot have two packages in one directory")
	} else {
		pkg = directory.Package
	}

	return
}
