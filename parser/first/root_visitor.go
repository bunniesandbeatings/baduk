package first

import (
	"go/ast"

	"github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"github.com/davecgh/go-spew/spew"
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
			receiverType := fieldListTypes(t.Recv)[0]
			params := fieldListTypes(t.Type.Params)
			returns := fieldListTypes(t.Type.Results)

			visitor.Package.AddMethod(t.Name.Name, visitor.Package.Name, visitor.Filename, receiverType, params, returns)
		}
	}
	return visitor
}

func fieldListTypes(fl *ast.FieldList) []architecture.Type {
	types := make([]architecture.Type, 0)
	if fl != nil {
		for _, field := range fl.List {
			types = append(types, typeOf(field.Type))
		}

	}
	return types
}

func typeOf(expr ast.Expr) architecture.Type {
	var t architecture.Type
	switch exprType := expr.(type) {

	case *ast.StarExpr:
		t = architecture.Type("*") + typeOf(exprType.X)

	case *ast.Ident:
		t = architecture.Type(exprType.Name)

	case *ast.ArrayType:
		len := exprType.Len
		if len == nil {
			t = architecture.Type("[]") + typeOf(exprType.Elt)
		} else {
			leng := len.(*ast.BasicLit).Value
			t = architecture.Type("["+leng+"]") + typeOf(exprType.Elt)
		}

	case *ast.Ellipsis:
		t = architecture.Type("...") + typeOf(exprType.Elt)

	case *ast.SelectorExpr:
		t = architecture.Type(exprType.X.(*ast.Ident).Name + "." + exprType.Sel.Name)

	case *ast.InterfaceType:
		spew.Printf("Not sure what to do with interface type %#v\n", expr)
		t = "interface-type-in-root"

	default:
		panic(spew.Sprintf("Cannot determine type of expression %#v", expr))
	}

	return t
}
