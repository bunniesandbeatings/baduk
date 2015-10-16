package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"
	"github.com/bunniesandbeatings/go-flavor-parser/parser"
)

func TestMatchSimpleMethodSet(t *testing.T) {
	pkg := parseMethodSetsPackage(t)
	checkMethodSetMatch(t, pkg, "Mixer", "ConcreteMixer")
}

func TestMatchMultiImplsMethodSet(t *testing.T) {
	pkg := parseMethodSetsPackage(t)
	checkMethodSetMatch(t, pkg, "Shape", "Circle")
	checkMethodSetMatch(t, pkg, "Shape", "Rectangle")
}

func TestMatchMultiIfacesMethodSet(t *testing.T) {
	pkg := parseMethodSetsPackage(t)
	checkMethodSetMatch(t, pkg, "Shape1D", "Circle")
	checkMethodSetMatch(t, pkg, "Shape1D", "Rectangle")
	checkMethodSetMatch(t, pkg, "Shape1D", "Rectangle")
	checkMethodSetMatch(t, pkg, "Shape2D", "Rectangle")
}

func checkMethodSetMatch(t *testing.T, pkg *architecture.Package, ifaceName, concreteTypeName string) {
	found := false
	for _, implementer := range findInterface(pkg.Interfaces, ifaceName).Implementers {
		if implementer == architecture.Type(concreteTypeName) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Matching failed")
	}
}

func findInterface(interfaces []*architecture.Interface, t string) *architecture.Interface {
	for _, iface := range interfaces {
		if iface.Name == t {
			return iface
		}
	}
	return nil
}

func parseMethodSetsPackage(t *testing.T) *architecture.Package {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	p := parser.NewParser(contexts.CreateBuildContext(contexts.CommandContext{
		GoPath: filepath.Join(pwd, "../fixtures/method_sets"),
	}))

	p.ParseImportSpec("test_method_sets")
	p.MatchMethodSets()
	arch := p.GetArchitecture()
	return arch.FindDirectory("test_method_sets").Package
}
