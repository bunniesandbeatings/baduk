package architecture

import (
	"go/ast"
	"go/types"
)

type Package struct {
	Name       string
	Structs    []*Struct
	Interfaces []*Interface
	Funcs      []*Func
	Methods    []*Method
	Info       *types.Info
	AstFiles   []*ast.File
}

func NewPackage(name string, astFiles []*ast.File) *Package {
	return &Package{
		Name: name,
		Info: &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
			Uses:  make(map[*ast.Ident]types.Object),
		},
		AstFiles: astFiles,
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

func (file *Package) AddMethod(name string, filename string, receiverType Type, parmtypes []Type, returnTypes []Type) (method *Method) {
	method = &Method{
		Func: Func{
			Name:        name,
			Filename:    filename,
			ParmTypes:   parmtypes,
			ReturnTypes: returnTypes,
		},
		ReceiverType: receiverType,
	}

	file.Methods = append(file.Methods, method)

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
