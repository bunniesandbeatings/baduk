package architecture

type Directory struct {
	Directories map[string]*Directory
	Package     string
	Files       map[string]*File
}

func newDirectory() *Directory {
	return &Directory{
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*File),
	}
}
