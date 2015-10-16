package first

import (
	"go/ast"

	"github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"github.com/davecgh/go-spew/spew"
)

type TypeSpecVisitor struct {
	Context
	TypeSpec *ast.TypeSpec
}

func NewTypeSpecVisitor(context Context, typeSpec *ast.TypeSpec) *TypeSpecVisitor {
	return &TypeSpecVisitor{
		context,
		typeSpec,
	}
}

func parseMethods(funcType *ast.FuncType) []architecture.Type {
	params := fieldListTypes(funcType.Params)
	returns := fieldListTypes(funcType.Results)

	result := []architecture.Type{"bunnies", "beatings"}
	for _, p := range params {
		result = append(result, p)
	}
	for _, r := range returns {
		result = append(result, r)
	}
	return result
}

func (visitor TypeSpecVisitor) Visit(node ast.Node) ast.Visitor {
	// TODO: Can types be private?

	switch t := node.(type) {
	case *ast.InterfaceType:
		methods := []*architecture.Method{}

		for i, field := range t.Methods.List {
			if len(field.Names) != 1 {
				panic(spew.Sprintf("Method %d of interface %s does not have one name: %#v", i, visitor.TypeSpec.Name.Name, field.Names))
			}

			var parmTypes, returnTypes []architecture.Type
			if t, ok := field.Type.(*ast.FuncType); ok {
				parmTypes = fieldListTypes(t.Params)
				returnTypes = fieldListTypes(t.Results)
			} else {
				panic(spew.Sprintf("Cannot determine type of interface field %#v", field))
			}

			method := &architecture.Method{
				Func: architecture.Func{
					Name:        field.Names[0].Name,
					Filename:    visitor.Filename,
					ParmTypes:   parmTypes,
					ReturnTypes: returnTypes,
				},
				ReceiverType: "",
			}
			methods = append(methods, method)
		}

		visitor.Package.AddInterface(visitor.TypeSpec.Name.Name, visitor.Filename, methods)
	case *ast.StructType:
		visitor.Package.AddStruct(visitor.TypeSpec.Name.Name, visitor.Filename)
	}
	return nil
}
