package main

import (
	"github.com/feynmaz/hex/internal/adapters/core/arithmetic"
	"github.com/feynmaz/hex/internal/ports"
)

func main() {
	// ports
	var core ports.Arithmetic

	core = arithmetic.NewAdapter()
	res, err := core.Addition(1, 3)
	if err != nil {
		panic(err)
	}
	println(res)
}
