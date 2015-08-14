package architecture

import (
	"go/token"
)

type File struct {
	PublicStructs    []string
	PublicInterfaces []string
	PublicFuncs      []string
}

type Directory struct {
	Directories map[string]*Directory
	Files       map[string]*File
	Fset        token.FileSet
}

func newDirectory() *Directory {
	return &Directory{
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*File),
	}
}
