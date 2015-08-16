package first

import (
	"go/ast"
)

type GenDeclVisitor struct {
	Context
}

func NewGenDeclVisitor(context Context) *GenDeclVisitor {
	return &GenDeclVisitor{context}
}

func (visitor GenDeclVisitor) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.TypeSpec:
		return NewTypeSpecVisitor(visitor.Context, t)
	}
	return nil
}
