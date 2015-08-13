package architecture

type File struct {
	PublicStructs    []string
	PublicInterfaces []string
	PublicFuncs      []string
}

type Directory struct {
	Directories map[string]*Directory
	Files       map[string]*File
}

func newDirectory() *Directory {
	return &Directory{
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*File),
	}
}
