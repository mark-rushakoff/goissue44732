package main

import (
	"github.com/mark-rushakoff/goissue44732/bar"
	"github.com/mark-rushakoff/goissue44732/foo"
)

func main() {
	b := &bar.Bar{}
	_ = foo.NewFoo()
	_ = b
}
