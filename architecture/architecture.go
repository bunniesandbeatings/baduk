package architecture

import "strings"

type Architecture struct {
	Root *Directory
}

func NewArchitecture() *Architecture {
	return &Architecture{
		Root: NewDirectory(),
	}
}

func (arch *Architecture) FindDirectory(path string) *Directory {
	pathSections := strings.Split(path, "/")

	currentNode := arch.Root

	for _, pathSection := range pathSections {
		if _, found := currentNode.Directories[pathSection]; !found {
			currentNode.Directories[pathSection] = NewDirectory()
		}

		currentNode = currentNode.Directories[pathSection]
	}

	return currentNode
}
