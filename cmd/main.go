package main

import "github.com/feynmaz/hex/internal/adapters/core/arithmetic"

func main() {
	arithmetic := arithmetic.NewAdapter()
	res, err := arithmetic.Addition(1,3)
	if err != nil {
		panic(err)
	}
	println(res)
}
