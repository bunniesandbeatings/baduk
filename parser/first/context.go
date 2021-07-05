package first

import (
	"go/token"

	"github.com/bunniesandbeatings/baduk/architecture"
)

type Context struct {
	Filename string
	Package  *architecture.Package
	Fset     *token.FileSet
}
