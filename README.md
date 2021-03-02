# goissue44732

```
$ gotip version
go version devel +2b50ab2aee Tue Mar 2 06:38:07 2021 +0000 darwin/amd64

$ gotip build main.go
# command-line-arguments
./main.go:11:4: internal compiler error: width not calculated: func()

goroutine 49 [running]:
runtime/debug.Stack(0x1b93380, 0xc00012a008, 0x0)
	/Users/mr/gotip/src/github.com/golang/go/src/runtime/debug/stack.go:24 +0x9f
cmd/compile/internal/base.FatalfAt(0xb04000000002, 0x1a5be61, 0x18, 0xc0007c9158, 0x1, 0x1)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/base/print.go:227 +0x1bc
cmd/compile/internal/base.Fatalf(...)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/base/print.go:196
cmd/compile/internal/types.CalcSize(0xc000433500)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/types/size.go:344 +0xf3
cmd/compile/internal/ssagen.TypeOK(0xc000433500, 0xc000433500)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:5243 +0x2f
cmd/compile/internal/ssagen.(*state).stmt(0xc0007a8100, 0x1ba5750, 0xc0007aa4b0)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:1530 +0x26c
cmd/compile/internal/ssagen.(*state).stmtList(0xc0007a8100, 0xc00079e060, 0x1, 0x1)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:1312 +0x68
cmd/compile/internal/ssagen.(*state).stmt(0xc0007a8100, 0x1ba59a8, 0xc0007a4140)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:1335 +0x1689
cmd/compile/internal/ssagen.(*state).stmtList(0xc0007a8100, 0xc0007a4180, 0x3, 0x4)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:1312 +0x68
cmd/compile/internal/ssagen.(*state).stmt(0xc0007a8100, 0x1ba5750, 0xc000411a90)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:1330 +0xf3
cmd/compile/internal/ssagen.(*state).stmtList(0xc0007a8100, 0xc0007a8000, 0xc, 0x10)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:1312 +0x68
cmd/compile/internal/ssagen.buildssa(0xc00040a160, 0x0, 0x0)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/ssa.go:565 +0x1645
cmd/compile/internal/ssagen.Compile(0xc00040a160, 0x0)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/ssagen/pgen.go:158 +0x5f
cmd/compile/internal/gc.compileFunctions.func2.1(0xc0007b0000, 0xc00040a160, 0xc0007a0028, 0xc000794090)
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/gc/compile.go:127 +0x65
created by cmd/compile/internal/gc.compileFunctions.func2
	/Users/mr/gotip/src/github.com/golang/go/src/cmd/compile/internal/gc/compile.go:125 +0x8e

$ go version
go version go1.15.3 darwin/amd64

$ go build
(okay, no output)
```
