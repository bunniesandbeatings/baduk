package first

import (
	target "github.com/bunniesandbeatings/go-flavor-parser/ast"
	"go/ast"
)

type RootVisitor struct {
	File *target.FileNode
}

func (visitor RootVisitor) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.GenDecl:
		return GenDeclVisitor{
			File: visitor.File,
		}
	case *ast.FuncDecl:
		// TODO: filter public only
		if t.Recv == nil {
			visitor.File.PublicFuncs = append(visitor.File.PublicFuncs, t.Name.Name)
		} else {
			// queue function with receiver
		}
	}
	return visitor
}
