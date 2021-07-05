package parser

import "github.com/bunniesandbeatings/baduk/architecture"

func (parser *Parser) MatchMethodSets() {
	for _, dir := range parser.arch.Root.Directories {
		for _, iface := range dir.Package.Interfaces {
			parser.matchInterface(iface)
		}
	}
}

func (parser *Parser) matchInterface(iface *architecture.Interface) {
	// Don't match empty interfaces.
	if len(iface.Methods) == 0 {
		return
	}

	ifaceMethod := iface.Methods[0]

	// all the methods of the interface must be methods of any type which implements the interface.
	// search for candidate concrete types and then see if their method sets match
	for _, mDir := range parser.arch.Root.Directories {
		for _, concreteMethod := range mDir.Package.Methods {
			if matchMethod(ifaceMethod, concreteMethod) && parser.matchMethodSet(iface, concreteMethod.Package, concreteMethod.ReceiverType) {
				implType := architecture.Type(concreteMethod.Package+".") + concreteMethod.ReceiverType
				iface.Implementers = add(iface.Implementers, implType)
			}
		}
	}
}

func (parser *Parser) matchMethodSet(iface *architecture.Interface, pkg string, receiverType architecture.Type) bool {
	for _, ifaceMethod := range iface.Methods {
		match := false
	scan:
		for _, mDir := range parser.arch.Root.Directories {
			for _, concreteMethod := range mDir.Package.Methods {
				if concreteMethod.Package == pkg &&
					concreteMethod.ReceiverType == receiverType &&
					matchMethod(ifaceMethod, concreteMethod) {
					match = true
					break scan
				}
			}
		}
		if !match {
			return false
		}
	}
	return true
}

func add(types []architecture.Type, t architecture.Type) []architecture.Type {
	for _, u := range types {
		if u == t {
			return types
		}
	}
	return append(types, t)
}

func matchMethod(m1, m2 *architecture.Method) bool {
	result := m1.Name == m2.Name &&
		matchTypeSlice(m1.ParmTypes, m2.ParmTypes) &&
		matchTypeSlice(m1.ReturnTypes, m2.ReturnTypes)
	return result
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
