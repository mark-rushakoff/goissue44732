package foo

type Foo struct {
	updatecb func()
}

func NewFoo() *Foo {
	f := &Foo{
		updatecb: nil,
	}
	return f
}
