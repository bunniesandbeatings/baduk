package ast

type FileNode struct {
	PublicStructs    []string
	PublicInterfaces []string
	PublicFuncs      []string
}

type PathNode struct {
	PathChildren map[string]*PathNode
	FileChildren map[string]*FileNode
}

func NewPathNode() *PathNode {
	return &PathNode{
		PathChildren: make(map[string]*PathNode),
	}
}


