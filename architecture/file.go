package architecture

type Package struct {
	Name       string
	Structs    []*Struct
	Interfaces []*Interface
	Funcs      []*Func
}

func NewPackage(name string) *Package {
	return &Package{
		Name: name,
	}
}

func (file *Package) AddFunc(name string, filename string) (function *Func) {
	function = &Func{
		Name:     name,
		Filename: filename,
	}

	file.Funcs = append(file.Funcs, function)

	return
}

func (file *Package) AddInterface(name string, filename string) (iface *Interface) {
	iface = &Interface{
		Name:     name,
		Filename: filename,
	}

	file.Interfaces = append(file.Interfaces, iface)

	return
}

func (file *Package) AddStruct(name string, filename string) (structure *Struct) {
	structure = &Struct{
		Name:     name,
		Filename: filename,
	}

	file.Structs = append(file.Structs, structure)

	return
}
