package first

import (
	target "github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"go/ast"
)

type TypeSpecVisitor struct {
	File     *target.File
	TypeSpec *ast.TypeSpec
}

func (visitor TypeSpecVisitor) Visit(node ast.Node) ast.Visitor {
	// TODO: Can types be private?

	switch node.(type) {
	case *ast.InterfaceType:
		visitor.File.AddInterface(visitor.TypeSpec.Name.Name)
	case *ast.StructType:
		visitor.File.AddStruct(visitor.TypeSpec.Name.Name)
	}
	return nil
}
