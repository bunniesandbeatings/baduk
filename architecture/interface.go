package architecture

type Interface struct {
	Name         string
	Filename     string
	Methods      []*Method
	Implementers []Type
}
