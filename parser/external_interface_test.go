package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bunniesandbeatings/go-flavor-parser/architecture"
	"github.com/bunniesandbeatings/go-flavor-parser/contexts"
	"github.com/bunniesandbeatings/go-flavor-parser/parser"
)

func TestMatchSimpleMethodSetToExternalInterface(t *testing.T) {
	pkg := parseMethodSetsAndExternalPackage(t)
	checkMethodSetMatchSize(t, pkg, "ExternalMixer", 1)
	checkMethodSetMatch(t, pkg, "ExternalMixer", "test_method_sets.ConcreteMixer")
}

func parseMethodSetsAndExternalPackage(t *testing.T) *architecture.Package {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	p := parser.NewParser(contexts.CreateBuildContext(contexts.CommandContext{
		GoPath: filepath.Join(pwd, "../fixtures/method_sets"),
	}))

	p.ParseImportSpec("test_method_sets")
	p.ParseImportSpec("external_interface")
	p.MatchMethodSets()
	arch := p.GetArchitecture()
	return arch.FindDirectory("external_interface").Package
}
