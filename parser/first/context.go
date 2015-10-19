package first

import (
	"go/token"

	"github.com/bunniesandbeatings/go-flavor-parser/architecture"
)

type Context struct {
	Filename string
	Package  *architecture.Package
	Fset     *token.FileSet
}
