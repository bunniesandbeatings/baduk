package first

import (
	"go/token"

	arch "github.com/bunniesandbeatings/go-flavor-parser/architecture"
)

type Context struct {
	Filename string
	Package  *arch.Package
	Fset     *token.FileSet
}
