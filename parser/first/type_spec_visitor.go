package first

import (
	target "github.com/bunniesandbeatings/go-flavor-parser/ast"
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
		visitor.File.PublicInterfaces = append(visitor.File.PublicInterfaces, visitor.TypeSpec.Name.Name)
	case *ast.StructType:
		visitor.File.PublicStructs = append(visitor.File.PublicStructs, visitor.TypeSpec.Name.Name)
	}
	return nil
}
