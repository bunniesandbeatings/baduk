package ast

type File struct {
	PublicStructs    []string
	PublicInterfaces []string
	PublicFuncs      []string
}

func NewFile() *File {
	return &File{}
}

type Directory struct {
	Directories map[string]*Directory
	Files map[string]*File
}

func NewDirectory() *Directory {
	return &Directory{
		Directories: make(map[string]*Directory),
		Files: make(map[string]*File),
	}
}

