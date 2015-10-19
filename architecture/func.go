package architecture

type Func struct {
	Name        string
	Package     string
	Filename    string
	ParmTypes   []Type
	ReturnTypes []Type
}
