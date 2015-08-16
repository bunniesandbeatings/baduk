package first

import (
	"go/ast"
)

type TypeSpecVisitor struct {
	Context
	TypeSpec *ast.TypeSpec
}

func NewTypeSpecVisitor(context Context, typeSpec *ast.TypeSpec) *TypeSpecVisitor {
	return &TypeSpecVisitor{
		context,
		typeSpec,
	}
}

func (visitor TypeSpecVisitor) Visit(node ast.Node) ast.Visitor {
	// TODO: Can types be private?

	switch node.(type) {
	case *ast.InterfaceType:
		visitor.Package.AddInterface(visitor.TypeSpec.Name.Name, visitor.Filename)
	case *ast.StructType:
		visitor.Package.AddStruct(visitor.TypeSpec.Name.Name, visitor.Filename)
	}
	return nil
}
