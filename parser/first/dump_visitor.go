package first

import (
	"fmt"
	"go/ast"
	"reflect"
)

type DumpVisitor struct {
}

func (visitor DumpVisitor) Visit(node ast.Node) ast.Visitor {
	fmt.Println(reflect.TypeOf(node))
	return visitor
}
