package test_types

import "code.google.com/p/gogoprotobuf/io"

type T int

func (T) Void() {
}

func (T) Scalar(int, string) (int64, uint32) {
	return 0, 0
}

func (T) Slice([]int) []string {
	return nil
}

func (T) Array([2]int) [3]string {
	ret := [3]string{"", "", ""}
	return ret
}

func (T) Pointer(*int) *string {
	return nil
}

func (T) Rich(a T, b []T, c *T) (*T, T, []T) {
	return c, a, b
}

func (T) Interface(i I) I {
	return nil
}

func (T) ImportedInterface(i io.Reader) io.Writer {
	return nil
}

func (T) DeeplyTyped([]*int, **[]**[2]**int) []***string {
	return []***string{}
}

func (T) Ellipsis(...int) {
}

type I interface {
	MethodOfInterface()
	AnotherMethodOfInterface()
}
