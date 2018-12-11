package main

import (
	"github.com/containous/dyngo/interp"
)

func main() {
	i := interp.New(interp.Opt{})
	i.Eval(`println("Hello")`)
}

// Output:
// Hello
