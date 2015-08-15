package first

import (
	"fmt"
	"go/ast"

	arch "github.com/bunniesandbeatings/go-flavor-parser/architecture"
)

type RootVisitor struct {
	File *arch.File
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
			visitor.File.AddFunc(t.Name.Name)
		} else {
			receiverType := getReceiverType(t.Recv)

			fmt.Printf("(rcvr %s) func %s()\n\n", receiverType, t.Name.Name)
		}
	}
	return visitor
}

func getReceiverType(receiver *ast.FieldList) string {
	switch receiverType := receiver.List[0].Type.(type) {

	case *ast.StarExpr:
		return receiverType.X.(*ast.Ident).Name

	case *ast.Ident:
		return receiverType.Name

	default:
		panic(fmt.Sprintf("Cannot Parse receiver: %v", receiver))
	}

	return ""
}
