package first

import (
	target "github.com/bunniesandbeatings/go-flavor-parser/ast"
	"go/ast"
)


type GenDeclVisitor struct {
	File *target.File
}

func (visitor GenDeclVisitor) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.TypeSpec:
		return TypeSpecVisitor{
			File: visitor.File,
			TypeSpec: t,
		}
	}
	return nil
}
