package test_method_sets

type ConcreteMixer int

func (ConcreteMixer) Mix() {}

type Mixer interface {
	Mix()
}
