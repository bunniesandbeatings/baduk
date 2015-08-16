package first

import (
	"fmt"
	"go/ast"
)

type RootVisitor struct {
	Context
}

func NewRootVisitor(context Context) *RootVisitor {
	return &RootVisitor{context}
}

func (visitor RootVisitor) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.GenDecl:
		return NewGenDeclVisitor(visitor.Context)
	case *ast.FuncDecl:
		// TODO: filter public only
		if t.Recv == nil {
			visitor.Package.AddFunc(t.Name.Name, visitor.Filename)
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
