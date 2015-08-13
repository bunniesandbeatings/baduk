package first

import (
	target "github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"go/ast"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type RootVisitor struct {
	File *target.File
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
			spew.Dump(t.Recv)
			fmt.Printf("(rcvr %v) func %s()\n\n", t.Recv.List[0].Type , t.Name.Name )
			// queue function with receiver
		}
	}
	return visitor
}
