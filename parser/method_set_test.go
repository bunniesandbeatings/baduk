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
	checkMethodSetMatch(t, pkg, "Mixer", "test_method_sets.ConcreteMixer")
}

func TestMatchMultiImplsMethodSet(t *testing.T) {
	pkg := parseMethodSetsPackage(t)
	checkMethodSetMatch(t, pkg, "Shape", "test_method_sets.Circle")
	checkMethodSetMatch(t, pkg, "Shape", "test_method_sets.Rectangle")
}

func TestMatchMultiIfacesMethodSet(t *testing.T) {
	pkg := parseMethodSetsPackage(t)
	checkMethodSetMatch(t, pkg, "Shape1D", "test_method_sets.Circle")
	checkMethodSetMatch(t, pkg, "Shape1D", "test_method_sets.Rectangle")
	checkMethodSetMatch(t, pkg, "Shape1D", "test_method_sets.Rectangle")
	checkMethodSetMatch(t, pkg, "Shape2D", "test_method_sets.Rectangle")
}

func checkMethodSetMatch(t *testing.T, pkg *architecture.Package, ifaceName, concreteTypeName string) {
	found := false

	iface := findInterface(pkg.Interfaces, ifaceName)
	if iface == nil {
		t.Errorf("Could not find interface %s in package %s", ifaceName, pkg.Name)
		t.FailNow()
	}

	for _, implementer := range iface.Implementers {
		if implementer == architecture.Type(concreteTypeName) {
			found = true
			break
		}
	}
	if !found {
		t.Logf("Implementatation %s of interface %s not found in implementers %#v", concreteTypeName, ifaceName, iface.Implementers)
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
