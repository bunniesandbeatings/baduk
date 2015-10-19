package parser

import "github.com/bunniesandbeatings/go-flavor-parser/architecture"

// Brute force until it becomes too slow...
func (parser *Parser) MatchMethodSets() {
	for _, dir := range parser.arch.Root.Directories {
		for _, iface := range dir.Package.Interfaces {
			parser.matchInterface(iface)
		}
	}
}

func (parser *Parser) matchInterface(iface *architecture.Interface) {
	for _, mDir := range parser.arch.Root.Directories {
		for _, method := range mDir.Package.Methods {
			for _, tDir := range parser.arch.Root.Directories {
				for _, concreteMethod := range tDir.Package.Methods {
					if matchMethod(method, concreteMethod) {
						implType := architecture.Type(concreteMethod.Package+".") + concreteMethod.ReceiverType
						iface.Implementers = append(iface.Implementers, implType)
					}
				}
			}
		}
	}
}

func matchMethod(m1, m2 *architecture.Method) bool {
	return m1.Name == m2.Name &&
		matchTypeSlice(m1.ParmTypes, m2.ParmTypes) &&
		matchTypeSlice(m1.ReturnTypes, m2.ReturnTypes)
}

func matchTypeSlice(s, t []architecture.Type) bool {
	if len(s) != len(t) {
		return false
	}
	for i, x := range s {
		// TODO: support interface types
		if x != t[i] {
			return false
		}
	}
	return true
}
