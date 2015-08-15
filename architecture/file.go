package architecture

type File struct {
	Structs []*Struct
	Interfaces []*Interface
	Funcs []*Func
}

func (file *File) AddFunc(name string) (function *Func) {
	function = &Func{
		Name: name,
	}

	file.Funcs = append(file.Funcs, function)

	return
}

func (file *File) AddInterface(name string) (iface *Interface) {
	iface = &Interface{
		Name: name,
	}

	file.Interfaces = append(file.Interfaces, iface)

	return
}

func (file *File) AddStruct(name string) (structure *Struct) {
	structure = &Struct{
		Name: name,
	}

	file.Structs = append(file.Structs, structure)

	return
}
