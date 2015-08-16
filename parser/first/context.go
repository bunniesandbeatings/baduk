package first

import (
	arch "github.com/bunniesandbeatings/go-flavor-parser/architecture"
)

type Context struct {
	Filename string
	Package  *arch.Package
}
